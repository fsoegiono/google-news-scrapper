package newsdata

type Source struct {
	Name string
	Logo string
}

type Data struct {
	Thumbnail string
	Title     string
	Url       string
	Source    Source
}
