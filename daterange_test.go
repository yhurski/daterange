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
