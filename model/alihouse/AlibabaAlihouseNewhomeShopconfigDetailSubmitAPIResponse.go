package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse 店铺配置信息接口 API返回值
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
type AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel
}

// AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel is 店铺配置信息接口 成功返回结果
type AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_detail_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeShopconfigDetailSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
