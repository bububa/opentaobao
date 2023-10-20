package cainiaocntec

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecCompassRpaExeResultsaveAPIResponse rpa执行结果回传 API返回值
// cainiao.cntec.compass.rpa.exe.resultsave
//
// rpa执行结果回传
type CainiaoCntecCompassRpaExeResultsaveAPIResponse struct {
	model.CommonResponse
	CainiaoCntecCompassRpaExeResultsaveAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCntecCompassRpaExeResultsaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCntecCompassRpaExeResultsaveAPIResponseModel).Reset()
}

// CainiaoCntecCompassRpaExeResultsaveAPIResponseModel is rpa执行结果回传 成功返回结果
type CainiaoCntecCompassRpaExeResultsaveAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_compass_rpa_exe_resultsave_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CainiaoCntecCompassRpaExeResultsaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCntecCompassRpaExeResultsaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCntecCompassRpaExeResultsaveAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCntecCompassRpaExeResultsaveAPIResponse)
	},
}

// GetCainiaoCntecCompassRpaExeResultsaveAPIResponse 从 sync.Pool 获取 CainiaoCntecCompassRpaExeResultsaveAPIResponse
func GetCainiaoCntecCompassRpaExeResultsaveAPIResponse() *CainiaoCntecCompassRpaExeResultsaveAPIResponse {
	return poolCainiaoCntecCompassRpaExeResultsaveAPIResponse.Get().(*CainiaoCntecCompassRpaExeResultsaveAPIResponse)
}

// ReleaseCainiaoCntecCompassRpaExeResultsaveAPIResponse 将 CainiaoCntecCompassRpaExeResultsaveAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCntecCompassRpaExeResultsaveAPIResponse(v *CainiaoCntecCompassRpaExeResultsaveAPIResponse) {
	v.Reset()
	poolCainiaoCntecCompassRpaExeResultsaveAPIResponse.Put(v)
}
