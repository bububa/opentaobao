package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderShipAPIRequest
闲鱼无忧购服务商发货 API请求
alibaba.idle.isv.order.ship

闲鱼无忧购业务入仓模式下服务商订单发货的接口 */
type AlibabaIdleIsvOrderShipAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId string
	// 物流公司
	_logisticsCompany string
	// 运单号
	_shipMailNo string
}

// NewAlibabaIdleIsvOrderShipRequest 初始化AlibabaIdleIsvOrderShipAPIRequest对象
func NewAlibabaIdleIsvOrderShipRequest() *AlibabaIdleIsvOrderShipAPIRequest {
	return &AlibabaIdleIsvOrderShipAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderShipAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderShipAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrderId Setter
// 订单号
func (r *AlibabaIdleIsvOrderShipAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r AlibabaIdleIsvOrderShipAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// Set is LogisticsCompany Setter
// 物流公司
func (r *AlibabaIdleIsvOrderShipAPIRequest) SetLogisticsCompany(_logisticsCompany string) error {
	r._logisticsCompany = _logisticsCompany
	r.Set("logistics_company", _logisticsCompany)
	return nil
}

// Get LogisticsCompany Getter
func (r AlibabaIdleIsvOrderShipAPIRequest) GetLogisticsCompany() string {
	return r._logisticsCompany
}

// Set is ShipMailNo Setter
// 运单号
func (r *AlibabaIdleIsvOrderShipAPIRequest) SetShipMailNo(_shipMailNo string) error {
	r._shipMailNo = _shipMailNo
	r.Set("ship_mail_no", _shipMailNo)
	return nil
}

// Get ShipMailNo Getter
func (r AlibabaIdleIsvOrderShipAPIRequest) GetShipMailNo() string {
	return r._shipMailNo
}
