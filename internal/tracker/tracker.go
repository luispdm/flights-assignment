package tracker

type Flight struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type tracker struct {
	fl     []Flight
	starts []string
	ends   []string
	d      differ
}

type differ interface {
	Diff(a, b []string) string
}

func New(fl []Flight, d differ) *tracker {
	t := &tracker{
		fl: fl,
		d:  d,
	}
	starts := make([]string, len(t.fl))
	ends := make([]string, len(t.fl))
	t.starts = starts
	t.ends = ends
	return t
}

func (t *tracker) Track() Flight {
	if len(t.fl) == 1 {
		return t.fl[0]
	}
	t.setSlices()
	return Flight{
		Start: t.d.Diff(t.starts, t.ends),
		End:   t.d.Diff(t.ends, t.starts),
	}
}

func (t *tracker) setSlices() {
	i := 0
	for _, v := range t.fl {
		t.starts[i] = v.Start
		t.ends[i] = v.End
		i++
	}
}
