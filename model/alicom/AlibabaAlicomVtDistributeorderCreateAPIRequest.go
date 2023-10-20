package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomVtDistributeorderCreateAPIRequest 通信业务外放下单 API请求
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
type AlibabaAlicomVtDistributeorderCreateAPIRequest struct {
	model.Params
	// 请求对象
	_distributeTradeMsgModel *DistributeTradeMsgModel
}

// NewAlibabaAlicomVtDistributeorderCreateRequest 初始化AlibabaAlicomVtDistributeorderCreateAPIRequest对象
func NewAlibabaAlicomVtDistributeorderCreateRequest() *AlibabaAlicomVtDistributeorderCreateAPIRequest {
	return &AlibabaAlicomVtDistributeorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomVtDistributeorderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.vt.distributeorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomVtDistributeorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomVtDistributeorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributeTradeMsgModel is DistributeTradeMsgModel Setter
// 请求对象
func (r *AlibabaAlicomVtDistributeorderCreateAPIRequest) SetDistributeTradeMsgModel(_distributeTradeMsgModel *DistributeTradeMsgModel) error {
	r._distributeTradeMsgModel = _distributeTradeMsgModel
	r.Set("distribute_trade_msg_model", _distributeTradeMsgModel)
	return nil
}

// GetDistributeTradeMsgModel DistributeTradeMsgModel Getter
func (r AlibabaAlicomVtDistributeorderCreateAPIRequest) GetDistributeTradeMsgModel() *DistributeTradeMsgModel {
	return r._distributeTradeMsgModel
}
