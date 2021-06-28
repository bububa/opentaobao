package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔订购公众号 APIResponse
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口
*/
type TaobaoJstSmsOfficialaccountOrderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_officialaccount_order_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统异常
    
    ResponseCode   string `json:"response_code,omitempty" xml:"