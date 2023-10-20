package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeshopconfigastorepreviewAPIResponse 天猫好房店铺装修-地址预览 API返回值
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
type AlibabaalihousenewhomeshopconfigastorepreviewAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeshopconfigastorepreviewAPIResponseModel
}

// AlibabaalihousenewhomeshopconfigastorepreviewAPIResponseModel is 天猫好房店铺装修-地址预览 成功返回结果
type AlibabaalihousenewhomeshopconfigastorepreviewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_astore_preview_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *AlibabaalihousenewhomeshopconfigastorepreviewResult `json:"result,omitempty" xml:"result,omitempty"`
}
