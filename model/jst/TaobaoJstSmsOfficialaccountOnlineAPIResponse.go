package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号上线 API返回值 
taobao.jst.sms.officialaccount.online

聚石塔公众号上线
*/
type TaobaoJstSmsOfficialaccountOnlineAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsOfficialaccountOnlineAPIResponseModel
}

// 聚石塔公众号上线 成功返回结果
type TaobaoJstSmsOfficialaccountOnlineAPIResponseModel struct {
    XMLName xml.Name `xml:"jst_sms_officialaccount_online_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统异常
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    // 请求id
    ResponseId   string `json:"response_id,omitempty" xml:"response_id,omitempty"`
    // 成功
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
    // 系统异常
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
