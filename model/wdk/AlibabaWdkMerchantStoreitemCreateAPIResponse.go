package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantStoreitemCreateAPIResponse 新建门店商品 API返回值
// alibaba.wdk.merchant.storeitem.create
//
// 新建门店商品
type AlibabaWdkMerchantStoreitemCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantStoreitemCreateAPIResponseModel
}

// AlibabaWdkMerchantStoreitemCreateAPIResponseModel is 新建门店商品 成功返回结果
type AlibabaWdkMerchantStoreitemCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_storeitem_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// errorCode
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// errorDesc
	Errordesc string `json:"errordesc,omitempty" xml:"errordesc,omitempty"`
}
