package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	googleInit()
}

func googleInit() {
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
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	// [END fs_add_data_1]

	// [START fs_add_data_2]
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
	// [END fs_add_data_2]

	// [START fs_get_all_users]
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	// [END fs_get_all_users]

	fmt.Println("hello")
	fmt.Println("asdf")
}

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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(listRoom())
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

// ===== Schedule
type Schedule struct {
	Name      string        `firestore:"name,omitempty" json:"name,omitempty"`
	MobileNo  string        `firestore:"mobile_no,omitempty" json:"mobile_no,omitempty"`
	Topic     string        `firestore:"topic,omitempty" json:"topic,omitempty"`
	Start     time.Time     `firestore:"time,omitempty" json:"time,omitempty"`
	StartUNIX int64         `firestore:"time_unix,omitempty" json:"time_unix,omitempty"`
	Duration  time.Duration `firestore:"duration,omitempty" json:"duration,omitempty"`
}

// ===== Schedule - ADD
func AddSchedule(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sched Schedule
	err := decoder.Decode(&sched)
	if err != nil {
		fmt.Println("error decoding", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sched.Start = time.Unix(sched.StartUNIX, 0)
	fmt.Println("parse from", sched.StartUNIX, "to", sched.Start)
	addSchedule(sched)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func addSchedule(sd Schedule) {
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

	client.Collection("schedules").Add(ctx, sd)
	if err != nil {
		log.Fatalf("Failed adding schedule: %+v, %v", sd, err)
	}
}

// ===== Schedule - List
func ListSchedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schds := listSchedule()
	l, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("error load location")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert time to +7
	for i, _ := range schds {
		fmt.Println("time:", schds[i].Start)
		schds[i].Start = schds[i].Start.In(l)
		fmt.Println("time in bkk:", schds[i].Start)
	}

	json.NewEncoder(w).Encode(schds)
}

func listSchedule() []Schedule {
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
	iter := client.Collection("schedules").Documents(ctx)
	scheds := []Schedule{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())

		var sd Schedule
		doc.DataTo(&sd)
		scheds = append(scheds, sd)
	}
	// [END fs_get_all_users]

	fmt.Printf("scheds: %+v", scheds)
	return scheds
}
