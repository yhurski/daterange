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
	inboundValue := midnightDate.After(dr.begin) && midnightDate.Before(dr.end)

	return boundValue || inboundValue
}

func (dr *DateRange) Entries() []time.Time {
	if len(dr.entries) == 0 {
		begin := dr.begin
		end := dr.end

		dr.entries = append(dr.entries, begin)
		for begin.Before(end) {
			begin = begin.Add(time.Hour * 24)
			dr.entries = append(dr.entries, begin)
		}
	}

	return dr.entries
}

func (dr *DateRange) Count() int {
	return len(dr.Entries())
}
