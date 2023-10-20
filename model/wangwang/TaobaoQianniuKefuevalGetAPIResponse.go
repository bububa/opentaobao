package wangwang

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuKefuevalGetAPIResponse 客服评价详情接口 API返回值
// taobao.qianniu.kefueval.get
//
// 获取买家对客服的服务评价
type TaobaoQianniuKefuevalGetAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuKefuevalGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuKefuevalGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuKefuevalGetAPIResponseModel).Reset()
}

// TaobaoQianniuKefuevalGetAPIResponseModel is 客服评价详情接口 成功返回结果
type TaobaoQianniuKefuevalGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_kefueval_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 评价明细
	StaffEvalDetails []EvalDetail `json:"staff_eval_details,omitempty" xml:"staff_eval_details>eval_detail,omitempty"`
	// 评价结果数
	ResultCount int64 `json:"result_count,omitempty" xml:"result_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuKefuevalGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StaffEvalDetails = m.StaffEvalDetails[:0]
	m.ResultCount = 0
}

var poolTaobaoQianniuKefuevalGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuKefuevalGetAPIResponse)
	},
}

// GetTaobaoQianniuKefuevalGetAPIResponse 从 sync.Pool 获取 TaobaoQianniuKefuevalGetAPIResponse
func GetTaobaoQianniuKefuevalGetAPIResponse() *TaobaoQianniuKefuevalGetAPIResponse {
	return poolTaobaoQianniuKefuevalGetAPIResponse.Get().(*TaobaoQianniuKefuevalGetAPIResponse)
}

// ReleaseTaobaoQianniuKefuevalGetAPIResponse 将 TaobaoQianniuKefuevalGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuKefuevalGetAPIResponse(v *TaobaoQianniuKefuevalGetAPIResponse) {
	v.Reset()
	poolTaobaoQianniuKefuevalGetAPIResponse.Put(v)
}
