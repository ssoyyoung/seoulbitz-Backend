package model

// DBinfo strutct
type DBinfo struct {
	Username string
	Password string
	Hostname string
	Port     string
	Database string
}

// Foodie  struct
type Foodie struct {
	Xpoint  string
	Ypoint  string
	Title   string
	Tag     string
	LikeCnt string
	Addr    string
	Insta   string
	Thumb   string
}

// Shopping struct
type Shopping struct {
	Xpoint  string
	Ypoint  string
	Title   string
	Tag     string
	LikeCnt string
	Addr    string
	Insta   string
	Thumb   string
}

// Subway struct
type Subway struct {
	Xpoint      string
	XpointWgs   string
	Ypoint      string
	YpointWgs   string
	StationNm   string
	StationCd   string
	LineNum     string
	FrCode      string
	CyberStCode string
}

// PlaceLatLng struct
type PlaceLatLng struct {
	Idx       int
	Title     string
	XpointWgs string
	YpointWgs string
}

// TwoPointDistance struct
type TwoPointDistance struct {
	Destination string
	Distance    float64
}
