package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillCloudprintNetprintPrintAPIResponse 网络打印机打印接口 API返回值
// cainiao.waybill.cloudprint.netprint.print
//
// 打印接口
type CainiaoWaybillCloudprintNetprintPrintAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillCloudprintNetprintPrintAPIResponseModel
}

// CainiaoWaybillCloudprintNetprintPrintAPIResponseModel is 网络打印机打印接口 成功返回结果
type CainiaoWaybillCloudprintNetprintPrintAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_cloudprint_netprint_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，0 为成功
	ServerErrorCode string `json:"server_error_code,omitempty" xml:"server_error_code,omitempty"`
	// 错误描述
	Describe string `json:"describe,omitempty" xml:"describe,omitempty"`
}
