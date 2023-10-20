package tbtrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradeorderskuupdateAPIResponse 更新交易的销售属性 API返回值
// taobao.trade.ordersku.update
//
// 只能更新发货前子订单的销售属性 &lt;br/&gt;只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 &lt;br/&gt;必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
type TaobaotradeorderskuupdateAPIResponse struct {
	model.CommonResponse
	TaobaotradeorderskuupdateAPIResponseModel
}

// TaobaotradeorderskuupdateAPIResponseModel is 更新交易的销售属性 成功返回结果
type TaobaotradeorderskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_ordersku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只返回oid和modified
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
}
