package daterange

import (
	"errors"
	"time"
)

type DateRange struct {
	begin    time.Time
	end      time.Time
	expanded []time.Time
}

func New(begin time.Time, end time.Time) (*DateRange, error) {
	if begin.After(end) {
		return nil, errors.New("first date is greater than the second")
	}

	return &DateRange{begin, end, make([]time.Time, 0, int(end.Sub(begin).Hours()/24))}, nil
}

func (dr *DateRange) In(date time.Time) bool {
	return date.After(dr.begin) && date.Before(dr.end)
}

func (dr *DateRange) Entries() []time.Time {
	if len(dr.expanded) == 0 {
		begin := dr.begin
		end := dr.end

		dr.expanded = append(dr.expanded, begin)
		for begin.Before(end) {
			begin = begin.Add(time.Hour * 24)
			dr.expanded = append(dr.expanded, begin)
		}
	}

	return dr.expanded
}

func (dr *DateRange) Count() int {
	return len(dr.Entries())
}
