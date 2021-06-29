package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新建门店商品 APIResponse
alibaba.wdk.merchant.storeitem.create

新建门店商品
*/
type AlibabaWdkMerchantStoreitemCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMerchantStoreitemCreateResponse
}

type AlibabaWdkMerchantStoreitemCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_merchant_storeitem_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    Suc   bool `json:"suc,omitempty" xml:"suc,omitempty"`

    
    // errorCode
    
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`

    
    // errorDesc
    
    Errordesc   string `json:"errordesc,omitempty" xml:"errordesc,omitempty"`

    
}
