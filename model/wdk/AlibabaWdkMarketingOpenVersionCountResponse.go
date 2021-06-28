package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
版本数量查询 APIResponse
alibaba.wdk.marketing.open.version.count

版本数量查询
*/
type AlibabaWdkMarketingOpenVersionCountAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_open_version_count_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"