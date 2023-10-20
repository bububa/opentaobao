package guoguo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoCpBackupAssigncourierAPIResponse CP兜底后指定接单的小件员 API返回值
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
type CainiaoGuoguoCpBackupAssigncourierAPIResponse struct {
	model.CommonResponse
	CainiaoGuoguoCpBackupAssigncourierAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGuoguoCpBackupAssigncourierAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGuoguoCpBackupAssigncourierAPIResponseModel).Reset()
}

// CainiaoGuoguoCpBackupAssigncourierAPIResponseModel is CP兜底后指定接单的小件员 成功返回结果
type CainiaoGuoguoCpBackupAssigncourierAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_cp_backup_assigncourier_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 错误信息
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGuoguoCpBackupAssigncourierAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StatusCode = ""
	m.StatusMessage = ""
	m.IsSuccess = false
}

var poolCainiaoGuoguoCpBackupAssigncourierAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGuoguoCpBackupAssigncourierAPIResponse)
	},
}

// GetCainiaoGuoguoCpBackupAssigncourierAPIResponse 从 sync.Pool 获取 CainiaoGuoguoCpBackupAssigncourierAPIResponse
func GetCainiaoGuoguoCpBackupAssigncourierAPIResponse() *CainiaoGuoguoCpBackupAssigncourierAPIResponse {
	return poolCainiaoGuoguoCpBackupAssigncourierAPIResponse.Get().(*CainiaoGuoguoCpBackupAssigncourierAPIResponse)
}

// ReleaseCainiaoGuoguoCpBackupAssigncourierAPIResponse 将 CainiaoGuoguoCpBackupAssigncourierAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGuoguoCpBackupAssigncourierAPIResponse(v *CainiaoGuoguoCpBackupAssigncourierAPIResponse) {
	v.Reset()
	poolCainiaoGuoguoCpBackupAssigncourierAPIResponse.Put(v)
}
