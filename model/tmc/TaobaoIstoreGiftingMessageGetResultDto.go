package tmc

// TaobaoIstoreGiftingMessageGetResultDTO 
type TaobaoIstoreGiftingMessageGetResultDTO struct {
    // id：消息id，receiverId：消息接收者，messageType：消息类型，messageDesc：消息描述，giftBoxId：礼盒id，attachInfo：消息附加参数，messageStatus：消息状态，sellerId：商家id，senderId：发送者
    ResultList   []string `json:"result_list,omitempty" xml:"result_list>string,omitempty"`
    // 附加信息
    BizExtMap   string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
