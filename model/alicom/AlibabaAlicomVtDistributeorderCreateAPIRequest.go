package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomvtdistributeordercreateAPIRequest 通信业务外放下单 API请求
// alibaba.alicom.vt.distributeorder.create
//
// 通信业务外放下单接口
type AlibabaalicomvtdistributeordercreateAPIRequest struct {
	model.Params
	// 请求对象
	_distributeTradeMsgModel *DistributeTradeMsgModel
}

// NewAlibabaalicomvtdistributeordercreateRequest 初始化AlibabaalicomvtdistributeordercreateAPIRequest对象
func NewAlibabaalicomvtdistributeordercreateRequest() *AlibabaalicomvtdistributeordercreateAPIRequest {
	return &AlibabaalicomvtdistributeordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomvtdistributeordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.vt.distributeorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomvtdistributeordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomvtdistributeordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributeTradeMsgModel is DistributeTradeMsgModel Setter
// 请求对象
func (r *AlibabaalicomvtdistributeordercreateAPIRequest) SetDistributeTradeMsgModel(_distributeTradeMsgModel *DistributeTradeMsgModel) error {
	r._distributeTradeMsgModel = _distributeTradeMsgModel
	r.Set("distribute_trade_msg_model", _distributeTradeMsgModel)
	return nil
}

// GetDistributeTradeMsgModel DistributeTradeMsgModel Getter
func (r AlibabaalicomvtdistributeordercreateAPIRequest) GetDistributeTradeMsgModel() *DistributeTradeMsgModel {
	return r._distributeTradeMsgModel
}
