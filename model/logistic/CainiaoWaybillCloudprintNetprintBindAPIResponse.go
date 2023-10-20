package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybillcloudprintnetprintbindAPIResponse 网络打印机绑定 API返回值
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
type CainiaowaybillcloudprintnetprintbindAPIResponse struct {
	model.CommonResponse
	CainiaowaybillcloudprintnetprintbindAPIResponseModel
}

// CainiaowaybillcloudprintnetprintbindAPIResponseModel is 网络打印机绑定 成功返回结果
type CainiaowaybillcloudprintnetprintbindAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_cloudprint_netprint_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *IsvResult `json:"result,omitempty" xml:"result,omitempty"`
}
