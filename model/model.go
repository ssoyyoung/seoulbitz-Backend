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
	Uniq    string
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
	Uniq    string
}

// Place struct
type Place struct {
	ShopType string  `json:"shop_type"`
	Xpoint   string  `json:"xpoint"`
	Ypoint   string  `json:"ypoint"`
	Title    string  `json:"title"`
	LikeCnt  string  `json:"like_cnt"`
	Addr     string  `json:"addr"`
	Insta    string  `json:"insta"`
	Thumb1   string  `json:"thumb1"`
	Thumb2   string  `json:"thumb2"`
	Distance float64 `json:"distance"`
	Uniq     string  `json:"uniq"`
}

// Subway struct
type Subway struct {
	Xpoint      string `json:"xpoint"`
	XpointWgs   string `json:"xpoint_wgs"`
	Ypoint      string `json:"ypoint"`
	YpointWgs   string `json:"ypoint_wgs"`
	StationNm   string `json:"station_nm"`
	StationCd   string `json:"station_cd"`
	LineNum     string `json:"line_num"`
	FrCode      string `json:"fr_code"`
	CyberStCode string `json:"cyber_st_code"`
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
	Title    string
	Distance float64
}

// NearPlace func
type NearPlace struct {
	Xpoint   string
	Ypoint   string
	Title    string
	Tag      string
	LikeCnt  string
	Addr     string
	Insta    string
	Thumb    string
	Distance float64
	Uniq     string
}
