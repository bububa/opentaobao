package xhotelitem

// TaobaoXhotelBaseinfoRoomGetResultSet 
type TaobaoXhotelBaseinfoRoomGetResultSet struct {

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 酒店基础信息
    
    XhotelBaseInfo   *XHotelInfoWithRoom `json:"xhotel_base_info,omitempty" xml:"xhotel_base_info,omitempty"`
    

}
