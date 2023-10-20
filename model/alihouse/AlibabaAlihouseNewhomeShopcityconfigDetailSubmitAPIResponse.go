package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse 城市配置信息接口 API返回值
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
type AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel
}

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel is 城市配置信息接口 成功返回结果
type AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopcityconfig_detail_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
