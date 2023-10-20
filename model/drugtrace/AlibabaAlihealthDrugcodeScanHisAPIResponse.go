package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeScanHisAPIResponse 企业查询扫码历史 API返回值
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
type AlibabaAlihealthDrugcodeScanHisAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeScanHisAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeScanHisAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeScanHisAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeScanHisAPIResponseModel is 企业查询扫码历史 成功返回结果
type AlibabaAlihealthDrugcodeScanHisAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeScanHisAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = m.Model[:0]
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAlibabaAlihealthDrugcodeScanHisAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeScanHisAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeScanHisAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeScanHisAPIResponse
func GetAlibabaAlihealthDrugcodeScanHisAPIResponse() *AlibabaAlihealthDrugcodeScanHisAPIResponse {
	return poolAlibabaAlihealthDrugcodeScanHisAPIResponse.Get().(*AlibabaAlihealthDrugcodeScanHisAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeScanHisAPIResponse 将 AlibabaAlihealthDrugcodeScanHisAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeScanHisAPIResponse(v *AlibabaAlihealthDrugcodeScanHisAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeScanHisAPIResponse.Put(v)
}
