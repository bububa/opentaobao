package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalModifyAPIResponse 修改审批单 API返回值
// alitrip.btrip.approval.modify
//
// 修改审批单
type AlitripBtripApprovalModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripApprovalModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripApprovalModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripApprovalModifyAPIResponseModel).Reset()
}

// AlitripBtripApprovalModifyAPIResponseModel is 修改审批单 成功返回结果
type AlitripBtripApprovalModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_approval_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripApprovalModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripApprovalModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripApprovalModifyAPIResponse)
	},
}

// GetAlitripBtripApprovalModifyAPIResponse 从 sync.Pool 获取 AlitripBtripApprovalModifyAPIResponse
func GetAlitripBtripApprovalModifyAPIResponse() *AlitripBtripApprovalModifyAPIResponse {
	return poolAlitripBtripApprovalModifyAPIResponse.Get().(*AlitripBtripApprovalModifyAPIResponse)
}

// ReleaseAlitripBtripApprovalModifyAPIResponse 将 AlitripBtripApprovalModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripApprovalModifyAPIResponse(v *AlitripBtripApprovalModifyAPIResponse) {
	v.Reset()
	poolAlitripBtripApprovalModifyAPIResponse.Put(v)
}
