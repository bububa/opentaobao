package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoRequisitionsGetAPIResponse 合作申请查询 API返回值
// taobao.fenxiao.requisitions.get
//
// 合作申请查询
type TaobaoFenxiaoRequisitionsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoRequisitionsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoRequisitionsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoRequisitionsGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoRequisitionsGetAPIResponseModel is 合作申请查询 成功返回结果
type TaobaoFenxiaoRequisitionsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_requisitions_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 合作申请
	Requisitions []Requisition `json:"requisitions,omitempty" xml:"requisitions>requisition,omitempty"`
	// 结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoRequisitionsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Requisitions = m.Requisitions[:0]
	m.TotalResults = 0
	m.IsSuccess = false
}

var poolTaobaoFenxiaoRequisitionsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoRequisitionsGetAPIResponse)
	},
}

// GetTaobaoFenxiaoRequisitionsGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoRequisitionsGetAPIResponse
func GetTaobaoFenxiaoRequisitionsGetAPIResponse() *TaobaoFenxiaoRequisitionsGetAPIResponse {
	return poolTaobaoFenxiaoRequisitionsGetAPIResponse.Get().(*TaobaoFenxiaoRequisitionsGetAPIResponse)
}

// ReleaseTaobaoFenxiaoRequisitionsGetAPIResponse 将 TaobaoFenxiaoRequisitionsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoRequisitionsGetAPIResponse(v *TaobaoFenxiaoRequisitionsGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoRequisitionsGetAPIResponse.Put(v)
}
