package cainiaohandover

// OpenServiceParam 
type OpenServiceParam struct {

    // DOOR_PICKUP:上门揽收；SELF_POST:自寄；SELF_SEND:自送；UNREACHABLE_RETURN:不可达退回；
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 不同物流服务的扩展信息
    
    Features   *Features `json:"features,omitempty" xml:"features,omitempty"`
    

}
