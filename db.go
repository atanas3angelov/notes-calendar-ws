package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// for security reasons read the user & pass within mongo connection string from
// untracked repo file
func getMongoDBconnectionString() string {
	mongodb_connect_string_file, err := os.Open("mongodb_atlas_connection_string.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mongodb_connect_string_file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	uri, err := io.ReadAll(mongodb_connect_string_file)

	return string(uri)
}

func getDBclient(uri string) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return client
}

func populateDBsample() {

	client := getDBclient(getMongoDBconnectionString())

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.
		Database("calendar_notes").
		Collection("notes")

	filter := bson.D{{Key: "year", Value: 2025}, {Key: "month", Value: 12}}

	var result Notes
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			var dayNotes = make(map[int]DayNotes)

			// Mock data, which should be collected from db instead
			dayNotes[3] = DayNotes{
				Summary: []string{"go"},
				Details: "A Tour of Go: variables & functions",
			}
			dayNotes[4] = DayNotes{
				Summary: []string{"go"},
				Details: "A Tour of Go: flow control - for",
			}
			dayNotes[5] = DayNotes{
				Summary: []string{"go", "python"},
				Details: "A Tour of Go: flow control - if, else, switch, defer\n" +
					"python - enumerate\n" + "c# closures\n" + "java enhanced for loop",
			}
			dayNotes[6] = DayNotes{
				Summary: []string{"go"},
				Details: "A Tour of Go: structs",
			}
			dayNotes[10] = DayNotes{
				Summary: []string{"go"},
				Details: "A Tour of Go: slices and maps",
			}
			dayNotes[11] = DayNotes{
				Summary: []string{"go"},
				Details: "A Tour of Go: exercise",
			}
			dayNotes[12] = DayNotes{
				Summary: []string{"go"},
			}
			dayNotes[17] = DayNotes{
				Summary: []string{"go"},
			}
			dayNotes[18] = DayNotes{
				Summary: []string{"go"},
				Details: "A Tour of Go: ",
			}
			dayNotes[20] = DayNotes{
				Summary: []string{"js", "css", "html"},
				Details: "creating a calendar with js",
			}
			dayNotes[25] = DayNotes{
				Summary: []string{"css"},
				Details: "css universal selector and combinators",
			}
			dayNotes[26] = DayNotes{
				Summary: []string{"js", "html"},
				Details: "figuring out how to display notes for the note calendar project",
			}

			monthNotes := Notes{ID: 202512, Year: 2025, Month: 12, Days: dayNotes}

			result, err := coll.InsertOne(context.TODO(), monthNotes)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

		} else {
			panic(err)
		}

	}
}

func getCalendarNotes(year, month int) map[int]DayNotes {

	client := getDBclient(getMongoDBconnectionString())

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.
		Database("calendar_notes").
		Collection("notes")

	var result Notes

	err := coll.FindOne(context.TODO(),
		bson.D{{Key: "year", Value: year}, {Key: "month", Value: month}}).
		Decode(&result)

	if err == mongo.ErrNoDocuments {
		return map[int]DayNotes{}
	}

	if err != nil {
		panic(err)
	}

	return result.Days
}

// Updates notes for a day (+year +month) or upserts a document (+year + month) for the day
// !!! Make sure the note summary and details are not empty at the same time!
func updateCalendarNote(year, month, day int, note DayNotes) {

	client := getDBclient(getMongoDBconnectionString())

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.
		Database("calendar_notes").
		Collection("notes")

	id, err := strconv.Atoi(fmt.Sprintf("%d%d", year, month))

	if err != nil {
		panic(err)
	}

	// if note is empty, throw panic - aborts changes
	if (len(note.Summary) == 0 ||
		(len(note.Summary) == 1 && len(note.Summary[0]) == 0)) &&
		len(note.Details) == 0 {

		panic(errors.New("empty note"))
	}

	// if note not empty and perform the update or do an upsert!
	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{
		bson.E{"$set", bson.M{"year": year}},   // can fix warnings using Key: & Value: as above
		bson.E{"$set", bson.M{"month": month}}, // but it would become harder to read
	} // for an upsert (if 1st note for a new month)

	// if summary field is not empty, replace with new value(s), else unset the document field
	if len(note.Summary) > 0 {
		update = append(update,
			bson.E{"$set", bson.M{fmt.Sprintf("days.%d.summary", day): note.Summary}})
	} else {
		update = append(update,
			bson.E{"$unset", bson.M{fmt.Sprintf("days.%d.summary", day): ""}})
	}

	// if details field is not empty, set the new value, else unset the document field
	if len(note.Details) > 0 {
		update = append(update,
			bson.E{"$set", bson.M{fmt.Sprintf("days.%d.details", day): note.Details}})
	} else {
		update = append(update,
			bson.E{"$unset", bson.M{fmt.Sprintf("days.%d.details", day): ""}})
	}

	opts := options.UpdateOne().SetUpsert(true)

	_, err = coll.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		panic(err)
	}
}

func deleteDayNotes(year, month, day int) {

	client := getDBclient(getMongoDBconnectionString())

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.
		Database("calendar_notes").
		Collection("notes")

	id, err := strconv.Atoi(fmt.Sprintf("%d%d", year, month))

	if err != nil {
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{"$unset", bson.M{fmt.Sprintf("days.%d", day): ""}}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}
}
