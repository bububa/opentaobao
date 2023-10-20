package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoCooperationGetAPIResponse 供应商或分销商获取合作关系信息 API返回值
// taobao.fenxiao.cooperation.get
//
// 获取供应商的合作关系信息
type TaobaoFenxiaoCooperationGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoCooperationGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoCooperationGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoCooperationGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoCooperationGetAPIResponseModel is 供应商或分销商获取合作关系信息 成功返回结果
type TaobaoFenxiaoCooperationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_cooperation_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 合作分销关系
	Cooperations []Cooperation `json:"cooperations,omitempty" xml:"cooperations>cooperation,omitempty"`
	// 结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoCooperationGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Cooperations = m.Cooperations[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoCooperationGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoCooperationGetAPIResponse)
	},
}

// GetTaobaoFenxiaoCooperationGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoCooperationGetAPIResponse
func GetTaobaoFenxiaoCooperationGetAPIResponse() *TaobaoFenxiaoCooperationGetAPIResponse {
	return poolTaobaoFenxiaoCooperationGetAPIResponse.Get().(*TaobaoFenxiaoCooperationGetAPIResponse)
}

// ReleaseTaobaoFenxiaoCooperationGetAPIResponse 将 TaobaoFenxiaoCooperationGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoCooperationGetAPIResponse(v *TaobaoFenxiaoCooperationGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoCooperationGetAPIResponse.Put(v)
}
