package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	ss := e[field]
	if len(ss) == 0 {
		return ""
	}
	return ss[0]
}
