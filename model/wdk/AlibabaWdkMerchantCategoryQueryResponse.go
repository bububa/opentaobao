package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三江erp对接类目查询接口 APIResponse
alibaba.wdk.merchant.category.query

三江erp对接类目查询接口
*/
type AlibabaWdkMerchantCategoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantCategoryQueryResponse
}

type AlibabaWdkMerchantCategoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_category_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    Suc   bool `json:"suc,omitempty" xml:"suc,omitempty"`

    
    // errorCode
    
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`

    
    // errorDesc
    
    Errordesc   string `json:"errordesc,omitempty" xml:"errordesc,omitempty"`

    
    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
