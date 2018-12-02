package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// ===== Schedule
type Schedule struct {
	ID        string        `firestore:"id,omitempty" json:"id,omitempty"`
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
	var scheds []Schedule
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
		sd.ID = doc.Ref.ID
		scheds = append(scheds, sd)
	}
	// [END fs_get_all_users]

	fmt.Printf("scheds: %+v", scheds)
	return scheds
}

// ===== Schedule - Filter
func FilterSchedule(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req struct {
		Begin int64 `firestore:"begin,omitempty" json:"begin,omitempty"`
		End   int64 `firestore:"end,omitempty" json:"end,omitempty"`
	}

	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println("error decoding", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	begin := time.Unix(req.Begin, 0)
	end := time.Unix(req.End, 0)

	fmt.Println("req.begin:", req.Begin)
	fmt.Println("req.end:", req.End)
	fmt.Println("begin:", begin)
	fmt.Println("end:", end)

	schedules := filterSchedule(begin, end)
	l, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("error load location")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert time to +7
	for i, _ := range schedules {
		fmt.Println("time:", schedules[i].Start)
		schedules[i].Start = schedules[i].Start.In(l)
		fmt.Println("time in bkk:", schedules[i].Start)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schedules)
}

func filterSchedule(begin, end time.Time) []Schedule {
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
	query := client.Collection("schedules").Where("time", ">=", begin).Where("time", "<=", end)
	iter := query.Documents(ctx)

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

	fmt.Printf("query: %+v", query)
	fmt.Printf("sched: %+v", scheds)
	return scheds
}
