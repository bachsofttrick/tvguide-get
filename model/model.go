package model

type ProgramDetails struct {
	Id            int64  `json:"id"`
	Season        int    `json:"seasonNumber"`
	EpisodeNumber int    `json:"episodeNumber"`
	EpisodeTitle  string `json:"episodeTitle"`
	Description   string `json:"description"`
	Rating        string `json:"tvRating"`
}

type Program struct {
	StartTime    int64 `json:"startTime"`
	UTCStartTime string
	ProgramId    int64  `json:"programId"`
	Title        string `json:"title"`
	Details      ProgramDetails
}

type ChannelInfo struct {
	Name     string `json:"fullName"`
	CallSign string `json:"name"`
	Number   string `json:"number"`
}

type Channel struct {
	Channel  ChannelInfo `json:"channel"`
	Schedule []Program   `json:"programSchedules"`
}

type tvGuideData struct {
	StartTime string    `json:"startTime"`
	Channels  []Channel `json:"items"`
}

type tvGuideDetailData struct {
	Item ProgramDetails `json:"item"`
}

type TVGuide struct {
	Data tvGuideData `json:"data"`
}

type TVGuideDetail struct {
	DetailData tvGuideDetailData `json:"data"`
}
