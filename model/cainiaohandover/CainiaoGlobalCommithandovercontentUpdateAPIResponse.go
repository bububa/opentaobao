package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalCommithandovercontentUpdateAPIResponse 修改已经提交的交接单 API返回值
// cainiao.global.commithandovercontent.update
//
// 修改已经提交的交接单
type CainiaoGlobalCommithandovercontentUpdateAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalCommithandovercontentUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalCommithandovercontentUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalCommithandovercontentUpdateAPIResponseModel).Reset()
}

// CainiaoGlobalCommithandovercontentUpdateAPIResponseModel is 修改已经提交的交接单 成功返回结果
type CainiaoGlobalCommithandovercontentUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_commithandovercontent_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	InternalErrorCode string `json:"internal_error_code,omitempty" xml:"internal_error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回对象
	Data *OpenHandoverContentUpdateResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功同result字段
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalCommithandovercontentUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InternalErrorCode = ""
	m.ErrorMsg = ""
	m.Data = nil
	m.IsSuccess = false
}

var poolCainiaoGlobalCommithandovercontentUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalCommithandovercontentUpdateAPIResponse)
	},
}

// GetCainiaoGlobalCommithandovercontentUpdateAPIResponse 从 sync.Pool 获取 CainiaoGlobalCommithandovercontentUpdateAPIResponse
func GetCainiaoGlobalCommithandovercontentUpdateAPIResponse() *CainiaoGlobalCommithandovercontentUpdateAPIResponse {
	return poolCainiaoGlobalCommithandovercontentUpdateAPIResponse.Get().(*CainiaoGlobalCommithandovercontentUpdateAPIResponse)
}

// ReleaseCainiaoGlobalCommithandovercontentUpdateAPIResponse 将 CainiaoGlobalCommithandovercontentUpdateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalCommithandovercontentUpdateAPIResponse(v *CainiaoGlobalCommithandovercontentUpdateAPIResponse) {
	v.Reset()
	poolCainiaoGlobalCommithandovercontentUpdateAPIResponse.Put(v)
}
