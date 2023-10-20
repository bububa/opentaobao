package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyAddAPIResponse 【商旅】isv添加审批单 API返回值
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
type AlitripBtripCorpopApplyAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplyAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopApplyAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopApplyAddAPIResponseModel).Reset()
}

// AlitripBtripCorpopApplyAddAPIResponseModel is 【商旅】isv添加审批单 成功返回结果
type AlitripBtripCorpopApplyAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参数
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopApplyAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopApplyAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopApplyAddAPIResponse)
	},
}

// GetAlitripBtripCorpopApplyAddAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopApplyAddAPIResponse
func GetAlitripBtripCorpopApplyAddAPIResponse() *AlitripBtripCorpopApplyAddAPIResponse {
	return poolAlitripBtripCorpopApplyAddAPIResponse.Get().(*AlitripBtripCorpopApplyAddAPIResponse)
}

// ReleaseAlitripBtripCorpopApplyAddAPIResponse 将 AlitripBtripCorpopApplyAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopApplyAddAPIResponse(v *AlitripBtripCorpopApplyAddAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopApplyAddAPIResponse.Put(v)
}
