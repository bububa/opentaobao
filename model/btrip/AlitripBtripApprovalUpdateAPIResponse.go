package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalUpdateAPIResponse 更新审批单 API返回值
// alitrip.btrip.approval.update
//
// 更新审批单
type AlitripBtripApprovalUpdateAPIResponse struct {
	model.CommonResponse
	AlitripBtripApprovalUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripApprovalUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripApprovalUpdateAPIResponseModel).Reset()
}

// AlitripBtripApprovalUpdateAPIResponseModel is 更新审批单 成功返回结果
type AlitripBtripApprovalUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_approval_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripApprovalUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripApprovalUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripApprovalUpdateAPIResponse)
	},
}

// GetAlitripBtripApprovalUpdateAPIResponse 从 sync.Pool 获取 AlitripBtripApprovalUpdateAPIResponse
func GetAlitripBtripApprovalUpdateAPIResponse() *AlitripBtripApprovalUpdateAPIResponse {
	return poolAlitripBtripApprovalUpdateAPIResponse.Get().(*AlitripBtripApprovalUpdateAPIResponse)
}

// ReleaseAlitripBtripApprovalUpdateAPIResponse 将 AlitripBtripApprovalUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripApprovalUpdateAPIResponse(v *AlitripBtripApprovalUpdateAPIResponse) {
	v.Reset()
	poolAlitripBtripApprovalUpdateAPIResponse.Put(v)
}
