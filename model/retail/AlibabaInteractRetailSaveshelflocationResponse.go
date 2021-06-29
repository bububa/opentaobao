package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存地理位置和货架关系 API返回值 
alibaba.interact.retail.saveshelflocation

保存地理位置和货架关系
*/
type AlibabaInteractRetailSaveshelflocationAPIResponse struct {
    model.CommonResponse
    AlibabaInteractRetailSaveshelflocationResponse
}

// 保存地理位置和货架关系 成功返回结果
type AlibabaInteractRetailSaveshelflocationResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_retail_saveshelflocation_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 保存结果
    Result   *AlibabaInteractRetailSaveshelflocationResult `json:"result,omitempty" xml:"result,omitempty"`
}
