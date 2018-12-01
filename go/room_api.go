package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Room struct {
	Building string `firestore:"building,omitempty" json:"building,omitempty"`
	Floor    string `firestore:"floor,omitempty" json:"floor,omitempty"`
	Type     string `firestore:"type,omitempty" json:"type,omitempty"`
	Capacity int    `firestore:"capacity,omitempty" json:"capacity,omitempty"`
	Name     string `firestore:"name,omitempty" json:"name,omitempty"`
}

// ===== ADD
func RoomAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rm Room
	err := decoder.Decode(&rm)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addRoom(rm)
	JSONResponse(w, http.StatusOK, nil)
}

func addRoom(rm Room) {
	ctx := context.Background()

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := "honghub-224111"

	// Get a Firestore client.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	// [END fs_initialize]

	// [START fs_add_data_1]
	_, _, err = client.Collection("rooms").Add(ctx, rm)
	if err != nil {
		log.Fatalf("Failed adding room: %+v, %v", rm, err)
	}
	// [END fs_add_data_1]
}

// ===== LIST
func ListRoom(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, http.StatusOK, listRoom())
}

func listRoom() []Room {
	ctx := context.Background()

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := "honghub-224111"

	// Get a Firestore client.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	// [END fs_initialize]

	// [START fs_get_all_users]
	iter := client.Collection("rooms").Documents(ctx)
	rooms := []Room{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())

		var rm Room
		doc.DataTo(&rm)
		rooms = append(rooms, rm)
	}
	// [END fs_get_all_users]

	fmt.Printf("rooms: %+v", rooms)
	return rooms
}
