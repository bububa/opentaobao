package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
备孕首页获取达人配置商品 APIResponse
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品
*/
type AlibabaAlihealthPregnancyProductGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyProductGetResponse
}

type AlibabaAlihealthPregnancyProductGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_product_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果集
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
