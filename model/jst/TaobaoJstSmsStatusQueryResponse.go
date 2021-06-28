package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号状态查询 APIResponse
taobao.jst.sms.status.query

聚石塔公众号状态查询
*/
type TaobaoJstSmsStatusQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_status_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *SmsResponse `json:"result,omitempty" xml:"