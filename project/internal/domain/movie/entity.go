package movie

type PaginationMovie struct {
	Movie       []Movie
	TotalResult string
}

type Movie struct {
	ImdbID string
	Type   string
	Title  string
	Year   string
	Poster string
}

type MovieDetail struct {
	ImdbID     string
	ImdbRating string
	ImdbVotes  string
	Type       string
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []MovieRating
	Metascore  string
}

type MovieRating struct {
	Source string
	Value  string
}
