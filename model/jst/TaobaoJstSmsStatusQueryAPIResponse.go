package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号状态查询 API返回值 
taobao.jst.sms.status.query

聚石塔公众号状态查询
*/
type TaobaoJstSmsStatusQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsStatusQueryAPIResponseModel
}

// 聚石塔公众号状态查询 成功返回结果
type TaobaoJstSmsStatusQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"jst_sms_status_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    Result   *SmsResponse `json:"result,omitempty" xml:"result,omitempty"`
}
