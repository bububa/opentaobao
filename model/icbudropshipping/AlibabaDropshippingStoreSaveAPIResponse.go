package icbudropshipping

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingStoreSaveAPIResponse 阿里巴巴dropshipping店铺数据保存接口 API返回值
// alibaba.dropshipping.store.save
//
// 阿里巴巴dropshipping店铺数据保存
type AlibabaDropshippingStoreSaveAPIResponse struct {
	model.CommonResponse
	AlibabaDropshippingStoreSaveAPIResponseModel
}

// AlibabaDropshippingStoreSaveAPIResponseModel is 阿里巴巴dropshipping店铺数据保存接口 成功返回结果
type AlibabaDropshippingStoreSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dropshipping_store_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//  is success
	ResultSuccess string `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
