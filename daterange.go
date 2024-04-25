package daterange

import (
	"errors"
	"time"
)

type DateRange struct {
	begin   time.Time
	end     time.Time
	entries []time.Time
}

func New(begin time.Time, end time.Time) (*DateRange, error) {
	byear, bmonth, bday := begin.Date()
	eyear, emonth, eday := end.Date()
	beginMidnight := time.Date(byear, bmonth, bday, 0, 0, 0, 0, time.UTC)
	endMidnight := time.Date(eyear, emonth, eday, 0, 0, 0, 0, time.UTC)

	if beginMidnight.After(endMidnight) {
		return nil, errors.New("first date is greater than the second")
	}

	newdr := DateRange{
		beginMidnight,
		endMidnight,
		make([]time.Time, 0, int(endMidnight.Sub(beginMidnight).Hours()/24)),
	}

	return &newdr, nil
}

func (dr *DateRange) In(date time.Time) bool {
	year, month, day := date.Date()
	midnightDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	boundValue := midnightDate == dr.begin || midnightDate == dr.end
	inclusiveValue := midnightDate.After(dr.begin) && midnightDate.Before(dr.end)

	return boundValue || inclusiveValue
}

func (dr *DateRange) Entries() []time.Time {
	if len(dr.entries) == 0 {
		begin := dr.begin
		end := dr.end

		dr.entries = append(dr.entries, begin)
		for begin != end {
			begin = begin.Add(time.Hour * 24)
			dr.entries = append(dr.entries, begin)
		}
	}

	return dr.entries
}

func (dr *DateRange) Count() int {
	return len(dr.Entries())
}

func (dr *DateRange) Eql(anotherDr DateRange) bool {
	return dr.begin == anotherDr.begin && dr.end == anotherDr.end
}

func (dr *DateRange) Begin() time.Time {
	return dr.begin
}

func (dr *DateRange) End() time.Time {
	return dr.end
}

func (dr *DateRange) First(n int) []time.Time {
	if n >= dr.Count() {
		return dr.Entries()
	}

	return dr.Entries()[:n]
}

func (dr *DateRange) Last(n int) []time.Time {
	numberOfEntries := dr.Count()
	if n >= numberOfEntries {
		return dr.Entries()
	}

	return dr.Entries()[numberOfEntries-n:]
}

// Returns true if anotherDr is between the begin and end of the range.
func (dr *DateRange) Cover(anotherDr DateRange) bool {
	beginInclusive := anotherDr.begin == dr.begin || anotherDr.begin.After(dr.begin)
	endInclusive := anotherDr.end == dr.end || anotherDr.end.Before(dr.end)

	return beginInclusive && endInclusive
}
