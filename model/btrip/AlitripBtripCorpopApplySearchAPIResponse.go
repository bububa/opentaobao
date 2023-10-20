package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplySearchAPIResponse 【商旅】搜索审批单列表 API返回值
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
type AlitripBtripCorpopApplySearchAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplySearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopApplySearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopApplySearchAPIResponseModel).Reset()
}

// AlitripBtripCorpopApplySearchAPIResponseModel is 【商旅】搜索审批单列表 成功返回结果
type AlitripBtripCorpopApplySearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopApplySearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopApplySearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopApplySearchAPIResponse)
	},
}

// GetAlitripBtripCorpopApplySearchAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopApplySearchAPIResponse
func GetAlitripBtripCorpopApplySearchAPIResponse() *AlitripBtripCorpopApplySearchAPIResponse {
	return poolAlitripBtripCorpopApplySearchAPIResponse.Get().(*AlitripBtripCorpopApplySearchAPIResponse)
}

// ReleaseAlitripBtripCorpopApplySearchAPIResponse 将 AlitripBtripCorpopApplySearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopApplySearchAPIResponse(v *AlitripBtripCorpopApplySearchAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopApplySearchAPIResponse.Put(v)
}
