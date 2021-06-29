package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商ISV闲鱼用户信息查询 API返回值 
alibaba.idle.isv.user.query

服务商ISV闲鱼用户信息查询
*/
type AlibabaIdleIsvUserQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvUserQueryResponse
}

// 服务商ISV闲鱼用户信息查询 成功返回结果
type AlibabaIdleIsvUserQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_user_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}
