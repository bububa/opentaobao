package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodescanhisAPIResponse 企业查询扫码历史 API返回值
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
type AlibabaalihealthdrugcodescanhisAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcodescanhisAPIResponseModel
}

// AlibabaalihealthdrugcodescanhisAPIResponseModel is 企业查询扫码历史 成功返回结果
type AlibabaalihealthdrugcodescanhisAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_scan_his_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Model []CodeScanDto `json:"model,omitempty" xml:"model>code_scan_dto,omitempty"`
	// 结果信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 结果编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
