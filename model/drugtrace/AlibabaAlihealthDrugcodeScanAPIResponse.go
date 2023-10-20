package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodescanAPIResponse 查询扫码信息 API返回值
// alibaba.alihealth.drugcode.scan
//
// 查询扫码信息
type AlibabaalihealthdrugcodescanAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcodescanAPIResponseModel
}

// AlibabaalihealthdrugcodescanAPIResponseModel is 查询扫码信息 成功返回结果
type AlibabaalihealthdrugcodescanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_scan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
