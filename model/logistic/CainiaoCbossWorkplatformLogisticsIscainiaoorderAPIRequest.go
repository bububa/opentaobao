package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest 根据交易单号判断是否为菜鸟发货订单 API请求
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
type CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId string
}

// NewCainiaocbossworkplatformlogisticsiscainiaoorderRequest 初始化CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest对象
func NewCainiaocbossworkplatformlogisticsiscainiaoorderRequest() *CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest {
	return &CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.logistics.iscainiaoorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// 交易单号
func (r *CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r CainiaocbossworkplatformlogisticsiscainiaoorderAPIRequest) GetTradeId() string {
	return r._tradeId
}
