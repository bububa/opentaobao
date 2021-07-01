package icbushowcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加橱窗商品 API返回值 
alibaba.scbp.showcase.addproduct

批量添加商品到橱窗
*/
type AlibabaScbpShowcaseAddproductAPIResponse struct {
    model.CommonResponse
    AlibabaScbpShowcaseAddproductAPIResponseModel
}

// 批量添加橱窗商品 成功返回结果
type AlibabaScbpShowcaseAddproductAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_showcase_addproduct_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否添加成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
