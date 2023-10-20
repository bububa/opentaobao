package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDiscountsGetAPIResponse 获取折扣信息 API返回值
// taobao.fenxiao.discounts.get
//
// 查询折扣信息
type TaobaoFenxiaoDiscountsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDiscountsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDiscountsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDiscountsGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoDiscountsGetAPIResponseModel is 获取折扣信息 成功返回结果
type TaobaoFenxiaoDiscountsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_discounts_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 折扣数据结构
	Discounts []Discount `json:"discounts,omitempty" xml:"discounts>discount,omitempty"`
	// 记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDiscountsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Discounts = m.Discounts[:0]
	m.TotalResults = 0
}

var poolTaobaoFenxiaoDiscountsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDiscountsGetAPIResponse)
	},
}

// GetTaobaoFenxiaoDiscountsGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDiscountsGetAPIResponse
func GetTaobaoFenxiaoDiscountsGetAPIResponse() *TaobaoFenxiaoDiscountsGetAPIResponse {
	return poolTaobaoFenxiaoDiscountsGetAPIResponse.Get().(*TaobaoFenxiaoDiscountsGetAPIResponse)
}

// ReleaseTaobaoFenxiaoDiscountsGetAPIResponse 将 TaobaoFenxiaoDiscountsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDiscountsGetAPIResponse(v *TaobaoFenxiaoDiscountsGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDiscountsGetAPIResponse.Put(v)
}
