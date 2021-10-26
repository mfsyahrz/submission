package omdb

type GetMovieResponse struct {
	Response    string
	Error       string
	Search      []Search
	TotalResult string `json:"totalResults"`
}

type Search struct {
	ImdbID string `json:"imdbID"`
	Type   string
	Title  string
	Year   string
	Poster string
}

type GetMovieDetailResponse struct {
	Response   string
	Error      string
	ImdbID     string `json:"imdbID"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
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
