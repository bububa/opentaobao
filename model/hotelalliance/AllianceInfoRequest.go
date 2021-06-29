package hotelalliance

// AllianceInfoRequest 
type AllianceInfoRequest struct {

    // 要查询的日期，格式yyyymmdd
    
    QueryDay   string `json:"query_day,omitempty" xml:"query_day,omitempty"`
    

    // 签约类型-0:融合；1:直签。
    
    SignType   int64 `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
    

}
