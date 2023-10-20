package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotradecarordergetAPIResponse 整车订单详情查询 API返回值
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
type TmallaliautotradecarordergetAPIResponse struct {
	model.CommonResponse
	TmallaliautotradecarordergetAPIResponseModel
}

// TmallaliautotradecarordergetAPIResponseModel is 整车订单详情查询 成功返回结果
type TmallaliautotradecarordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_trade_car_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
