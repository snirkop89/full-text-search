package fulltextsearch

type Index map[string][]int

func (i Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := i[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice
				continue
			}
			i[token] = append(ids, doc.ID)
		}
	}
}

func intersection(a, b []int) []int {
	maxLen := max(len(a), len(b))
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (i Index) Search(query string) []int {
	var r []int
	for _, token := range analyze(query) {
		if ids, ok := i[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// Token doesn't exist
			return nil
		}
	}

	return r
}
