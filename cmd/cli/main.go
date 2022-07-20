package main

import (
	"log"
	"sync"
)

type flight struct {
	Start string
	End   string
}

var oneFlight = []flight{
	{
		Start: "ATL",
		End:   "EWR",
	},
}

// start: BGY; end: AKL
var flights = []flight{
	{
		Start: "BCN",
		End:   "PSC",
	},
	{
		Start: "JFK",
		End:   "AAL",
	},
	{
		Start: "FCO",
		End:   "BCN",
	},
	{
		Start: "GSO",
		End:   "IND",
	},
	{
		Start: "SFO",
		End:   "ATL",
	},
	{
		Start: "AAL",
		End:   "HEL",
	},
	{
		Start: "PSC",
		End:   "BLQ",
	},
	{
		Start: "IND",
		End:   "EWR",
	},
	{
		Start: "BGY",
		End:   "RAR",
	},
	{
		Start: "BJZ",
		End:   "AKL",
	},
	{
		Start: "AUH",
		End:   "FCO",
	},
	{
		Start: "HEL",
		End:   "CAK",
	},
	{
		Start: "RAR",
		End:   "AUH",
	},
	{
		Start: "CAK",
		End:   "BJZ",
	},
	{
		Start: "ATL",
		End:   "GSO",
	},
	{
		Start: "CHI",
		End:   "JFK",
	},
	{
		Start: "BLQ",
		End:   "MAD",
	},
	{
		Start: "EWR",
		End:   "CHI",
	},
	{
		Start: "MAD",
		End:   "SFO",
	},
}

func main() {
	log.Println(sliceDiff(flights))
	log.Println(mapDiff(flights))
	log.Println(doubleFor(flights, map[string]bool{}, map[string]bool{}))
	log.Println(doubleForCh(flights, &sync.Map{}, &sync.Map{}))
}

func sliceDiff(flights []flight) (string, string) {
	starts := make([]string, len(flights))
	ends := make([]string, len(flights))
	i := 0
	for _, v := range flights {
		starts[i] = v.Start
		ends[i] = v.End
		i++
	}
	return sDiff(starts, ends), sDiff(ends, starts)
}

func sDiff(a, b []string) string {
	bMap := make(map[string]bool, len(b))
	for _, v := range b {
		bMap[v] = true
	}
	diff := ""
	for _, v := range a {
		if !bMap[v] {
			diff = v
			break
		}
	}
	return diff
}

func mapDiff(flights []flight) (string, string) {
	starts := map[string]bool{}
	ends := map[string]bool{}
	for _, v := range flights {
		starts[v.Start] = true
		ends[v.End] = true
	}
	return mDiff(starts, ends), mDiff(ends, starts)
}

func mDiff(a, b map[string]bool) string {
	sSlice := make([]string, len(a))
	diff := ""
	i := 0
	for k := range a {
		sSlice[i] = k
		i++
	}
	for j := range sSlice {
		if !b[sSlice[j]] {
			diff = sSlice[j]
		}
	}
	return diff
}

func doubleFor(flights []flight, s, e map[string]bool) (string, string) {
	start := ""
	end := ""
	for _, v := range flights {
		for _, w := range flights {
			if start != "" || (!s[v.Start] && v.Start == w.End) {
				s[v.Start] = true
			}
			if end != "" || (!e[v.End] && v.End == w.Start) {
				e[v.End] = true
			}
		}
		if !s[v.Start] {
			start = v.Start
		}
		if !e[v.End] {
			end = v.End
		}
		if start != "" && end != "" {
			break
		}
	}
	return start, end
}

func doubleForCh(flights []flight, s, e *sync.Map) (string, string) {
	start := make(chan string, 1)
	end := make(chan string, 1)
	wg := sync.WaitGroup{}

	for _, v := range flights {
		wg.Add(1)
		go func(v flight) {
			defer wg.Done()
			for _, w := range flights {
				_, nSFound := s.Load(v.Start)
				_, nEFound := e.Load(v.End)
				if !nSFound && v.Start == w.End {
					s.Store(v.Start, true)
				}
				if !nEFound && v.End == w.Start {
					e.Store(v.End, true)
				}
			}
			if _, ok := s.Load(v.Start); !ok {
				start <- v.Start
			}
			if _, ok := e.Load(v.End); !ok {
				end <- v.End
			}
		}(v)
	}

	go func(st, en chan string) {
		wg.Wait()
		close(st)
		close(en)
	}(start, end)

	return <-start, <-end
}
