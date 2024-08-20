package model

type Program struct {
	StartTime    int64 `json:"startTime"`
	UTCStartTime string
	ProgramId    int64  `json:"programId"`
	Title        string `json:"title"`
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

type TVGuide struct {
	Data tvGuideData `json:"data"`
}
