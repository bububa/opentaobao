package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest 根据交易单号判断是否为菜鸟发货订单 API请求
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId string
}

// NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest 初始化CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest对象
func NewCainiaoCbossWorkplatformLogisticsIscainiaoorderRequest() *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest {
	return &CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.logistics.iscainiaoorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeId is TradeId Setter
// 交易单号
func (r *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest) GetTradeId() string {
	return r._tradeId
}
