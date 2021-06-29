package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
数据同步版本号申请 API返回值 
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请
*/
type AlibabaWdkMarketingOpenVersionApplyAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingOpenVersionApplyResponse
}

// 数据同步版本号申请 成功返回结果
type AlibabaWdkMarketingOpenVersionApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_open_version_apply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 版本号申请结果
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
