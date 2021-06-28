package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
数据同步版本号申请 APIResponse
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请
*/
type AlibabaWdkMarketingOpenVersionApplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_open_version_apply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 版本号申请结果
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"