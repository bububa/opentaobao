package refund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundRefuseAPIResponse 卖家拒绝退款 API返回值
// taobao.refund.refuse
//
// 卖家拒绝单笔退款（包含退款和退款退货）交易，要求如下：&lt;br/&gt;1. 传入的refund_id和相应的tid, oid必须匹配&lt;br/&gt;2. 如果一笔订单只有一笔子订单，则tid必须与oid相同&lt;br/&gt;3. 只有卖家才能执行拒绝退款操作&lt;br/&gt;4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单
type TaobaoRefundRefuseAPIResponse struct {
	model.CommonResponse
	TaobaoRefundRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRefundRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRefundRefuseAPIResponseModel).Reset()
}

// TaobaoRefundRefuseAPIResponseModel is 卖家拒绝退款 成功返回结果
type TaobaoRefundRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段
	Refund *Refund `json:"refund,omitempty" xml:"refund,omitempty"`
	// 拒绝退款操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRefundRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Refund = nil
	m.IsSuccess = false
}

var poolTaobaoRefundRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRefundRefuseAPIResponse)
	},
}

// GetTaobaoRefundRefuseAPIResponse 从 sync.Pool 获取 TaobaoRefundRefuseAPIResponse
func GetTaobaoRefundRefuseAPIResponse() *TaobaoRefundRefuseAPIResponse {
	return poolTaobaoRefundRefuseAPIResponse.Get().(*TaobaoRefundRefuseAPIResponse)
}

// ReleaseTaobaoRefundRefuseAPIResponse 将 TaobaoRefundRefuseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRefundRefuseAPIResponse(v *TaobaoRefundRefuseAPIResponse) {
	v.Reset()
	poolTaobaoRefundRefuseAPIResponse.Put(v)
}
