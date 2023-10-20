package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeshopconfigastoresyncAPIResponse 天猫好房店铺装修-Astore上翻 API返回值
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
type AlibabaalihousenewhomeshopconfigastoresyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeshopconfigastoresyncAPIResponseModel
}

// AlibabaalihousenewhomeshopconfigastoresyncAPIResponseModel is 天猫好房店铺装修-Astore上翻 成功返回结果
type AlibabaalihousenewhomeshopconfigastoresyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_astore_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *AlibabaalihousenewhomeshopconfigastoresyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
