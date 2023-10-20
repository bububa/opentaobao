package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopApplyModifyAPIResponse 【商旅】修改出差审批单（行程） API返回值
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
type AlitripBtripCorpopApplyModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopApplyModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopApplyModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopApplyModifyAPIResponseModel).Reset()
}

// AlitripBtripCorpopApplyModifyAPIResponseModel is 【商旅】修改出差审批单（行程） 成功返回结果
type AlitripBtripCorpopApplyModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopApplyModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopApplyModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopApplyModifyAPIResponse)
	},
}

// GetAlitripBtripCorpopApplyModifyAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopApplyModifyAPIResponse
func GetAlitripBtripCorpopApplyModifyAPIResponse() *AlitripBtripCorpopApplyModifyAPIResponse {
	return poolAlitripBtripCorpopApplyModifyAPIResponse.Get().(*AlitripBtripCorpopApplyModifyAPIResponse)
}

// ReleaseAlitripBtripCorpopApplyModifyAPIResponse 将 AlitripBtripCorpopApplyModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopApplyModifyAPIResponse(v *AlitripBtripCorpopApplyModifyAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopApplyModifyAPIResponse.Put(v)
}
