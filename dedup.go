package dedup

func preinitMap(l int) map[string]bool {
	return make(map[string]bool, l)
}

func justMap() map[string]bool {
	return make(map[string]bool)
}

func dedup(ss []string, m map[string]bool) []string {
	dss := []string{}
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			dss = append(dss, s)
			m[s] = true
		}
	}
	return dss
}
