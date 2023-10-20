package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillcloudprintnetprintverifycodeAPIResponse 打印验证码 API返回值
// cainiao.waybill.cloudprint.netprint.verifycode
//
// 打印获取验证码
type CainiaowaybillcloudprintnetprintverifycodeAPIResponse struct {
	model.CommonResponse
	CainiaowaybillcloudprintnetprintverifycodeAPIResponseModel
}

// CainiaowaybillcloudprintnetprintverifycodeAPIResponseModel is 打印验证码 成功返回结果
type CainiaowaybillcloudprintnetprintverifycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_cloudprint_netprint_verifycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *IsvResult `json:"result,omitempty" xml:"result,omitempty"`
}
