package daterange

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	t.Run("Begin date is greater than end date", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2024, 3, 1, 10, 10, 10, 10, time.UTC)
		_, err := New(t1, t2)

		if err != nil {
			t.Errorf("Failed: %s", err)
		}
	})

	t.Run("Begin date is equal to end date", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		// t2 := time.Date(2024, 3, 1, 10, 10, 10, 10, time.UTC)
		_, err := New(t1, t2)

		if err != nil {
			t.Errorf("Failed: %s", err)
		}
	})

	t.Run("Begin date shouldn't be greater than end date", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2023, 1, 1, 10, 10, 10, 10, time.UTC)
		// t2 := time.Date(2024, 3, 1, 10, 10, 10, 10, time.UTC)
		_, err := New(t1, t2)

		if err == nil {
			t.Errorf("Failed: %s", err)
		}
	})
}

func TestIn(t *testing.T) {
	t.Run("Date in interval should be included", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2024, 3, 1, 10, 10, 10, 10, time.UTC)
		checkDate := time.Date(2024, 1, 5, 10, 10, 10, 10, time.UTC)
		dr, _ := New(t1, t2)

		if !dr.In(checkDate) {
			t.Errorf("Not in the inverval: %s", checkDate)
		}

	})
}

func TestEntries(t *testing.T) {
	t.Run("The same begin and end dates should have one entry", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		dr, _ := New(t1, t2)
		entries := dr.Entries()

		if len(entries) != 1 || entries[0] != time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) {
			t.Error("Entries does not contain one item")
		}
	})

	t.Run("The same begin and end dates should have one entry", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2024, 1, 15, 22, 12, 15, 10, time.UTC)
		dr, _ := New(t1, t2)
		entries := dr.Entries()
		accurateEntries := 15

		if len(entries) != accurateEntries {
			t.Errorf("Entries does not contain %d item", accurateEntries)
		}
	})
}
