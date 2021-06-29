package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
备孕首页获取达人配置商品 API返回值 
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品
*/
type AlibabaAlihealthPregnancyProductGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyProductGetResponse
}

// 备孕首页获取达人配置商品 成功返回结果
type AlibabaAlihealthPregnancyProductGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_product_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
