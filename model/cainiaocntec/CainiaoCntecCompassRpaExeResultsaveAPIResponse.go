package cainiaocntec

import (
	"encoding/xml"

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

// CainiaoCntecCompassRpaExeResultsaveAPIResponseModel is rpa执行结果回传 成功返回结果
type CainiaoCntecCompassRpaExeResultsaveAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_compass_rpa_exe_resultsave_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CainiaoCntecCompassRpaExeResultsaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
