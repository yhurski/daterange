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

func TestCount(t *testing.T) {
	t.Run("Count should return corrent number of items", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		dr, _ := New(t1, t2)
		items := dr.Count()
		accurateItems := 366 + 15

		if items != accurateItems {
			t.Errorf("Entries should contain %d item but contains %d items", accurateItems, items)
		}
	})
}

func TestEql(t *testing.T) {
	t.Run("Two ranges with the same begin and end dates should be the same", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		dr, _ := New(t1, t2)

		anotherT1 := time.Date(2024, 1, 1, 10, 12, 30, 22, time.UTC)
		anotherT2 := time.Date(2025, 1, 15, 20, 00, 1, 2, time.UTC)
		anotherDr, _ := New(anotherT1, anotherT2)

		if !dr.Eql(*anotherDr) {
			t.Error("Date should be the same")
		}
	})

	t.Run("Two ranges with the same begin but different end dates shouldn't be the same", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 12, 30, 22, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		dr, _ := New(t1, t2)

		anotherT1 := time.Date(2024, 1, 1, 10, 12, 30, 22, time.UTC)
		anotherT2 := time.Date(2025, 1, 14, 20, 00, 1, 2, time.UTC)
		anotherDr, _ := New(anotherT1, anotherT2)

		if dr.Eql(*anotherDr) {
			t.Error("Date shouldn't be the same")
		}
	})

	t.Run("Two ranges with the same end but different begin dates shouldn't be the same", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		dr, _ := New(t1, t2)

		anotherT1 := time.Date(2024, 1, 2, 10, 12, 30, 22, time.UTC)
		anotherT2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		anotherDr, _ := New(anotherT1, anotherT2)

		if dr.Eql(*anotherDr) {
			t.Error("Date shouldn't be the same")
		}
	})
}

func TestBegin(t *testing.T) {
	t.Run("Begin date should be equal to the first date of range", func(t *testing.T) {
		t1 := time.Date(2024, 9, 28, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		begin := time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC)
		dr, _ := New(t1, t2)

		if dr.Begin() != begin {
			t.Errorf("Begin is %s but should be %s\n", dr.Begin(), begin)
		}
	})
}

func TestEnd(t *testing.T) {
	t.Run("End date should be equal to the last date of range", func(t *testing.T) {
		t1 := time.Date(2024, 9, 28, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		end := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
		dr, _ := New(t1, t2)

		if dr.End() != end {
			t.Errorf("End is %s but should be %s\n", dr.End(), end)
		}
	})
}

func TestFirst(t *testing.T) {
	t.Run("First should return first n items", func(t *testing.T) {
		t1 := time.Date(2024, 1, 1, 10, 10, 10, 10, time.UTC)
		t2 := time.Date(2025, 1, 15, 22, 12, 15, 10, time.UTC)
		dr, _ := New(t1, t2)
		items := dr.First(3)
		datesToReturn := [...]time.Time{
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		}

		for i, date := range datesToReturn {
			if date != items[i] {
				t.Errorf("Date should be %s but is %s", date, items[i])
			}
		}
	})
}
