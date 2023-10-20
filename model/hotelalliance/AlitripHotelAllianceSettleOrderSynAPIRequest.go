package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelalliancesettleordersynAPIRequest 菲住联盟分账成功订单同步 API请求
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
type AlitriphotelalliancesettleordersynAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *AllianceSettleOrderInfo
}

// NewAlitriphotelalliancesettleordersynRequest 初始化AlitriphotelalliancesettleordersynAPIRequest对象
func NewAlitriphotelalliancesettleordersynRequest() *AlitriphotelalliancesettleordersynAPIRequest {
	return &AlitriphotelalliancesettleordersynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelalliancesettleordersynAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.alliance.settle.order.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelalliancesettleordersynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelalliancesettleordersynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 订单信息
func (r *AlitriphotelalliancesettleordersynAPIRequest) SetOrderInfo(_orderInfo *AllianceSettleOrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlitriphotelalliancesettleordersynAPIRequest) GetOrderInfo() *AllianceSettleOrderInfo {
	return r._orderInfo
}
