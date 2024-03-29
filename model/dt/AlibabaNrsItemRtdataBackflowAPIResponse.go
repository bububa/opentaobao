package dt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNrsItemRtdataBackflowAPIResponse RT竞价数据回流 API返回值
// alibaba.nrs.item.rtdata.backflow
//
// 回流竞品价格数据，用与后续OCR识别价签数据，做精确化数据纠正
type AlibabaNrsItemRtdataBackflowAPIResponse struct {
	model.CommonResponse
	AlibabaNrsItemRtdataBackflowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNrsItemRtdataBackflowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNrsItemRtdataBackflowAPIResponseModel).Reset()
}

// AlibabaNrsItemRtdataBackflowAPIResponseModel is RT竞价数据回流 成功返回结果
type AlibabaNrsItemRtdataBackflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nrs_item_rtdata_backflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *NrsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNrsItemRtdataBackflowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaNrsItemRtdataBackflowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNrsItemRtdataBackflowAPIResponse)
	},
}

// GetAlibabaNrsItemRtdataBackflowAPIResponse 从 sync.Pool 获取 AlibabaNrsItemRtdataBackflowAPIResponse
func GetAlibabaNrsItemRtdataBackflowAPIResponse() *AlibabaNrsItemRtdataBackflowAPIResponse {
	return poolAlibabaNrsItemRtdataBackflowAPIResponse.Get().(*AlibabaNrsItemRtdataBackflowAPIResponse)
}

// ReleaseAlibabaNrsItemRtdataBackflowAPIResponse 将 AlibabaNrsItemRtdataBackflowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNrsItemRtdataBackflowAPIResponse(v *AlibabaNrsItemRtdataBackflowAPIResponse) {
	v.Reset()
	poolAlibabaNrsItemRtdataBackflowAPIResponse.Put(v)
}
