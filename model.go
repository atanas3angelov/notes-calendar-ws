package main

type DayNotes struct {
	Summary []string `bson:"summary,omitempty" json:"summary"`
	Details string   `bson:"details,omitempty" json:"details"`
}

type Notes struct {
	ID    int              `bson:"_id"`
	Year  int              `bson:"year"`
	Month int              `bson:"month"`
	Days  map[int]DayNotes `bson:"days,omitempty"`
}
