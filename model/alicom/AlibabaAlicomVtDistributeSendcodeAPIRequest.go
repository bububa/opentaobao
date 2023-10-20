package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomvtdistributesendcodeAPIRequest 通信业务外放发送验证码 API请求
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
type AlibabaalicomvtdistributesendcodeAPIRequest struct {
	model.Params
	// 请求对象
	_distributeTradeMsgModel *DistributeTradeMsgModel
}

// NewAlibabaalicomvtdistributesendcodeRequest 初始化AlibabaalicomvtdistributesendcodeAPIRequest对象
func NewAlibabaalicomvtdistributesendcodeRequest() *AlibabaalicomvtdistributesendcodeAPIRequest {
	return &AlibabaalicomvtdistributesendcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalicomvtdistributesendcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.vt.distribute.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalicomvtdistributesendcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalicomvtdistributesendcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributeTradeMsgModel is DistributeTradeMsgModel Setter
// 请求对象
func (r *AlibabaalicomvtdistributesendcodeAPIRequest) SetDistributeTradeMsgModel(_distributeTradeMsgModel *DistributeTradeMsgModel) error {
	r._distributeTradeMsgModel = _distributeTradeMsgModel
	r.Set("distribute_trade_msg_model", _distributeTradeMsgModel)
	return nil
}

// GetDistributeTradeMsgModel DistributeTradeMsgModel Getter
func (r AlibabaalicomvtdistributesendcodeAPIRequest) GetDistributeTradeMsgModel() *DistributeTradeMsgModel {
	return r._distributeTradeMsgModel
}
