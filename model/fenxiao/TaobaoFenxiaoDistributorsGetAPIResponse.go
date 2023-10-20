package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorsGetAPIResponse 获取分销商信息 API返回值
// taobao.fenxiao.distributors.get
//
// 查询和当前登录供应商有合作关系的分销商的信息
type TaobaoFenxiaoDistributorsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDistributorsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDistributorsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDistributorsGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoDistributorsGetAPIResponseModel is 获取分销商信息 成功返回结果
type TaobaoFenxiaoDistributorsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_distributors_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分销商详细信息
	Distributors []Distributor `json:"distributors,omitempty" xml:"distributors>distributor,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDistributorsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Distributors = m.Distributors[:0]
}

var poolTaobaoFenxiaoDistributorsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDistributorsGetAPIResponse)
	},
}

// GetTaobaoFenxiaoDistributorsGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDistributorsGetAPIResponse
func GetTaobaoFenxiaoDistributorsGetAPIResponse() *TaobaoFenxiaoDistributorsGetAPIResponse {
	return poolTaobaoFenxiaoDistributorsGetAPIResponse.Get().(*TaobaoFenxiaoDistributorsGetAPIResponse)
}

// ReleaseTaobaoFenxiaoDistributorsGetAPIResponse 将 TaobaoFenxiaoDistributorsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDistributorsGetAPIResponse(v *TaobaoFenxiaoDistributorsGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDistributorsGetAPIResponse.Put(v)
}
