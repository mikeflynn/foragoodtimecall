package main

import ()

type Contest struct {
	ID        int64  `json:id`
	Title     string `json:title`
	StartsOn  string `json:starts_on`
	EndsOn    string `json:ends_on`
	WinnerID  int64  `json:winner_id`
	CreatedOn string `json:created_on`
}

func NewContest() (*Contest, error) {
	return &Contest{}, nil
}

func (this *Contest) Load() bool {
	return true
}

func (this *Contest) Save() bool {
	return true
}
