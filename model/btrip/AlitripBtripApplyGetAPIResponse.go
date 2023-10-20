package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApplyGetAPIResponse 获取单个审批单 API返回值
// alitrip.btrip.apply.get
//
// 获取单个审批单的详情数据
type AlitripBtripApplyGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripApplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripApplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripApplyGetAPIResponseModel).Reset()
}

// AlitripBtripApplyGetAPIResponseModel is 获取单个审批单 成功返回结果
type AlitripBtripApplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_apply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripApplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripApplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripApplyGetAPIResponse)
	},
}

// GetAlitripBtripApplyGetAPIResponse 从 sync.Pool 获取 AlitripBtripApplyGetAPIResponse
func GetAlitripBtripApplyGetAPIResponse() *AlitripBtripApplyGetAPIResponse {
	return poolAlitripBtripApplyGetAPIResponse.Get().(*AlitripBtripApplyGetAPIResponse)
}

// ReleaseAlitripBtripApplyGetAPIResponse 将 AlitripBtripApplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripApplyGetAPIResponse(v *AlitripBtripApplyGetAPIResponse) {
	v.Reset()
	poolAlitripBtripApplyGetAPIResponse.Put(v)
}
