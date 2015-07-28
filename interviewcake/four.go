package interviewcake

import "sort"

type Meeting struct {
	StartTime int
	EndTime   int
}

type SortableMeetings []Meeting

func (s SortableMeetings) Len() int      { return len(s) }
func (s SortableMeetings) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortableMeetings) Less(i, j int) bool {
	return s[i].StartTime < s[j].StartTime
}

func CondenseMeetingTimes(allMeetings []Meeting) []Meeting {
	if len(allMeetings) == 0 {
		return nil
	}
	sort.Sort(SortableMeetings(allMeetings))
	var condensed []Meeting
	current := allMeetings[0]
	for i := 1; i < len(allMeetings); i++ {
		m := allMeetings[i]
		if m.StartTime > current.EndTime {
			condensed = append(condensed, current)
			current = m
			continue
		}
		if m.EndTime > current.EndTime {
			current.EndTime = m.EndTime
		}
	}
	condensed = append(condensed, current)
	return condensed
}
