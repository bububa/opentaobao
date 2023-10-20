package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponse 城市配置信息接口 API返回值
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
type AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponseModel
}

// AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponseModel is 城市配置信息接口 成功返回结果
type AlibabaalihousenewhomeshopcityconfigdetailsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopcityconfig_detail_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeshopcityconfigdetailsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
