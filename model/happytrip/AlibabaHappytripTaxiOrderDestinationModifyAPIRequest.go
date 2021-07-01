package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderDestinationModifyAPIRequest
修改目的地 API请求
alibaba.happytrip.taxi.order.destination.modify

通知ISV修改订单信息 */
type AlibabaHappytripTaxiOrderDestinationModifyAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 目的地经度
	_tlng string
	// 目的地纬度
	_tlat string
	// 目的地名称(最多50个字)
	_endName string
	// 目的地详细地址(最多100个字)
	_endAddress string
}

// New
