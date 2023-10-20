package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopDepartSyncAPIResponse 外部部门同步 API返回值
// alitrip.btrip.corpop.depart.sync
//
// 同步外部平台部门信息至商旅内部
type AlitripBtripCorpopDepartSyncAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopDepartSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopDepartSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopDepartSyncAPIResponseModel).Reset()
}

// AlitripBtripCorpopDepartSyncAPIResponseModel is 外部部门同步 成功返回结果
type AlitripBtripCorpopDepartSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_depart_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 导入失败原因以及导入失败部门信息
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 正确
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopDepartSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Module = ""
	m.ResultCode = 0
	m.SuccessFlag = false
}

var poolAlitripBtripCorpopDepartSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopDepartSyncAPIResponse)
	},
}

// GetAlitripBtripCorpopDepartSyncAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopDepartSyncAPIResponse
func GetAlitripBtripCorpopDepartSyncAPIResponse() *AlitripBtripCorpopDepartSyncAPIResponse {
	return poolAlitripBtripCorpopDepartSyncAPIResponse.Get().(*AlitripBtripCorpopDepartSyncAPIResponse)
}

// ReleaseAlitripBtripCorpopDepartSyncAPIResponse 将 AlitripBtripCorpopDepartSyncAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopDepartSyncAPIResponse(v *AlitripBtripCorpopDepartSyncAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopDepartSyncAPIResponse.Put(v)
}
