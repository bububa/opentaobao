package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradedrugordersgetAPIResponse 阿里健康获取某一药店全部订单 API返回值
// taobao.trade.drug.orders.get
//
// 阿里健康获取某一药店全部订单
type TaobaotradedrugordersgetAPIResponse struct {
	model.CommonResponse
	TaobaotradedrugordersgetAPIResponseModel
}

// TaobaotradedrugordersgetAPIResponseModel is 阿里健康获取某一药店全部订单 成功返回结果
type TaobaotradedrugordersgetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的订单列表对象
	Result *TaobaotradedrugordersgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
