package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripApplySearchAPIResponse 搜索审批单 API返回值
// alitrip.btrip.apply.search
//
// 外部企业调用获取本企业审批单列表数据
type AlitripBtripApplySearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripApplySearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripApplySearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripApplySearchAPIResponseModel).Reset()
}

// AlitripBtripApplySearchAPIResponseModel is 搜索审批单 成功返回结果
type AlitripBtripApplySearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_apply_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BtriphomeResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripApplySearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripApplySearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripApplySearchAPIResponse)
	},
}

// GetAlitripBtripApplySearchAPIResponse 从 sync.Pool 获取 AlitripBtripApplySearchAPIResponse
func GetAlitripBtripApplySearchAPIResponse() *AlitripBtripApplySearchAPIResponse {
	return poolAlitripBtripApplySearchAPIResponse.Get().(*AlitripBtripApplySearchAPIResponse)
}

// ReleaseAlitripBtripApplySearchAPIResponse 将 AlitripBtripApplySearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripApplySearchAPIResponse(v *AlitripBtripApplySearchAPIResponse) {
	v.Reset()
	poolAlitripBtripApplySearchAPIResponse.Put(v)
}
