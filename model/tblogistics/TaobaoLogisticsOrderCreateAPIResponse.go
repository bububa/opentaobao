package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOrderCreateAPIResponse 创建物流订单 API返回值
// taobao.logistics.order.create
//
// 用户调用该接口可以创建物流订单。目前仅支持手工订单的创建，创建完毕后默认自动使用“自己联系”的方式发货并且初始状态为”已发货“。也可以通过可选参数选择是否发货以及何种方式进行发货。
type TaobaoLogisticsOrderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsOrderCreateAPIResponseModel).Reset()
}

// TaobaoLogisticsOrderCreateAPIResponseModel is 创建物流订单 成功返回结果
type TaobaoLogisticsOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝物流订单交易号，如返回-1则表示错误。如果在新建订单时传入trade_id,此处会返回此id，如果未传入trade_id，此处会返回淘宝物流分配的交易号码。
	TradeId int64 `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeId = 0
}

var poolTaobaoLogisticsOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsOrderCreateAPIResponse)
	},
}

// GetTaobaoLogisticsOrderCreateAPIResponse 从 sync.Pool 获取 TaobaoLogisticsOrderCreateAPIResponse
func GetTaobaoLogisticsOrderCreateAPIResponse() *TaobaoLogisticsOrderCreateAPIResponse {
	return poolTaobaoLogisticsOrderCreateAPIResponse.Get().(*TaobaoLogisticsOrderCreateAPIResponse)
}

// ReleaseTaobaoLogisticsOrderCreateAPIResponse 将 TaobaoLogisticsOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsOrderCreateAPIResponse(v *TaobaoLogisticsOrderCreateAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsOrderCreateAPIResponse.Put(v)
}
