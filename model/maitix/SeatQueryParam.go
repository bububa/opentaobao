package maitix

// SeatQueryParam 
type SeatQueryParam struct {
    // 有条件必传-看台信息,如果传了下面的seat_info信息可以不传这个,否则必传
    Stands   []StandQueryParam `json:"stands,omitempty" xml:"stands>stand_query_param,omitempty"`
    // 必传-场次ID
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 必传-项目ID
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
    // 必传--token
    Token   string `json:"token,omitempty" xml:"token,omitempty"`
    // 必传-会话ID,必须和请求token接口传入的request_id一致.不是top接口返回的requestId
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 有条件必传-城市ID,如果传了下面的seat_info信息可以不传这个
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    // 可选,选座信息,只需要把回调的seatInfoV2对应的json数据urlDecode后直接赋值给这个字段就行,
    SeatInfo   string `json:"seat_info,omitempty" xml:"seat_info,omitempty"`
}
