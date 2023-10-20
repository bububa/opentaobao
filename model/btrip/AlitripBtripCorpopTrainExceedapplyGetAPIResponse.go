package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopTrainExceedapplyGetAPIResponse 商旅火车票第三方超标审批单搜索接口 API返回值
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
type AlitripBtripCorpopTrainExceedapplyGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopTrainExceedapplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopTrainExceedapplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopTrainExceedapplyGetAPIResponseModel).Reset()
}

// AlitripBtripCorpopTrainExceedapplyGetAPIResponseModel is 商旅火车票第三方超标审批单搜索接口 成功返回结果
type AlitripBtripCorpopTrainExceedapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_train_exceedapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopTrainExceedapplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopTrainExceedapplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopTrainExceedapplyGetAPIResponse)
	},
}

// GetAlitripBtripCorpopTrainExceedapplyGetAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopTrainExceedapplyGetAPIResponse
func GetAlitripBtripCorpopTrainExceedapplyGetAPIResponse() *AlitripBtripCorpopTrainExceedapplyGetAPIResponse {
	return poolAlitripBtripCorpopTrainExceedapplyGetAPIResponse.Get().(*AlitripBtripCorpopTrainExceedapplyGetAPIResponse)
}

// ReleaseAlitripBtripCorpopTrainExceedapplyGetAPIResponse 将 AlitripBtripCorpopTrainExceedapplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopTrainExceedapplyGetAPIResponse(v *AlitripBtripCorpopTrainExceedapplyGetAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopTrainExceedapplyGetAPIResponse.Put(v)
}
