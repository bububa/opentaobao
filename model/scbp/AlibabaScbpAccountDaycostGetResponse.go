package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询今日消耗 API返回值 
alibaba.scbp.account.daycost.get

查询今日消耗
*/
type AlibabaScbpAccountDaycostGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAccountDaycostGetResponse
}

// 查询今日消耗 成功返回结果
type AlibabaScbpAccountDaycostGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_account_daycost_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回今日消耗，单位元，两位小数
    DayCost   string `json:"day_cost,omitempty" xml:"day_cost,omitempty"`
}
