package main

type Params struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Organization string `json:"organization"`
	Space        string `json:"space"`
	SkipCleanup  bool   `json:"skip_cleanup"`
}
