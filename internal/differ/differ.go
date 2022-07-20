package differ

type differ struct {
}

func New() *differ {
	return &differ{}
}

func (d *differ) Diff(a, b []string) string {
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
