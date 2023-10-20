package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomvtdistributequeryprotocolAPIRequest 通信业务外放协议查询 API请求
// alibaba.alicom.vt.distribute.queryprotocol
//
// 通信业务外放协议查询
type AlibabaalicomvtdistributequeryprotocolAPIRequest struct {
	model.Params
	// 请求参数
	_distributeTradeMsgModel *DistributeTradeMsgModel
}

// NewAlibabaalicomvtdistributequeryprotocolRequest 初始化AlibabaalicomvtdistributequeryprotocolAPIRequest对象
func NewAlibabaalicomvtdistributequeryprotocolRequest() *AlibabaalicomvtdistributequeryprotocolAPIRequest {
	return &AlibabaalicomvtdistributequeryprotocolAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomvtdistributequeryprotocolAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.vt.distribute.queryprotocol"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomvtdistributequeryprotocolAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomvtdistributequeryprotocolAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributeTradeMsgModel is DistributeTradeMsgModel Setter
// 请求参数
func (r *AlibabaalicomvtdistributequeryprotocolAPIRequest) SetDistributeTradeMsgModel(_distributeTradeMsgModel *DistributeTradeMsgModel) error {
	r._distributeTradeMsgModel = _distributeTradeMsgModel
	r.Set("distribute_trade_msg_model", _distributeTradeMsgModel)
	return nil
}

// GetDistributeTradeMsgModel DistributeTradeMsgModel Getter
func (r AlibabaalicomvtdistributequeryprotocolAPIRequest) GetDistributeTradeMsgModel() *DistributeTradeMsgModel {
	return r._distributeTradeMsgModel
}
