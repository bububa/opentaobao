package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApprovalNewAPIResponse 新建审批单 API返回值
// alitrip.btrip.approval.new
//
// 用户新建审批单
type AlitripBtripApprovalNewAPIResponse struct {
	model.CommonResponse
	AlitripBtripApprovalNewAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripApprovalNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripApprovalNewAPIResponseModel).Reset()
}

// AlitripBtripApprovalNewAPIResponseModel is 新建审批单 成功返回结果
type AlitripBtripApprovalNewAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_approval_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripApprovalNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripApprovalNewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripApprovalNewAPIResponse)
	},
}

// GetAlitripBtripApprovalNewAPIResponse 从 sync.Pool 获取 AlitripBtripApprovalNewAPIResponse
func GetAlitripBtripApprovalNewAPIResponse() *AlitripBtripApprovalNewAPIResponse {
	return poolAlitripBtripApprovalNewAPIResponse.Get().(*AlitripBtripApprovalNewAPIResponse)
}

// ReleaseAlitripBtripApprovalNewAPIResponse 将 AlitripBtripApprovalNewAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripApprovalNewAPIResponse(v *AlitripBtripApprovalNewAPIResponse) {
	v.Reset()
	poolAlitripBtripApprovalNewAPIResponse.Put(v)
}
