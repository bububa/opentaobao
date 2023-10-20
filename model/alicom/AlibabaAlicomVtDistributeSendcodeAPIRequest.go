package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomVtDistributeSendcodeAPIRequest 通信业务外放发送验证码 API请求
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
type AlibabaAlicomVtDistributeSendcodeAPIRequest struct {
	model.Params
	// 请求对象
	_distributeTradeMsgModel *DistributeTradeMsgModel
}

// NewAlibabaAlicomVtDistributeSendcodeRequest 初始化AlibabaAlicomVtDistributeSendcodeAPIRequest对象
func NewAlibabaAlicomVtDistributeSendcodeRequest() *AlibabaAlicomVtDistributeSendcodeAPIRequest {
	return &AlibabaAlicomVtDistributeSendcodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomVtDistributeSendcodeAPIRequest) Reset() {
	r._distributeTradeMsgModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomVtDistributeSendcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.vt.distribute.sendcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomVtDistributeSendcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlicomVtDistributeSendcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributeTradeMsgModel is DistributeTradeMsgModel Setter
// 请求对象
func (r *AlibabaAlicomVtDistributeSendcodeAPIRequest) SetDistributeTradeMsgModel(_distributeTradeMsgModel *DistributeTradeMsgModel) error {
	r._distributeTradeMsgModel = _distributeTradeMsgModel
	r.Set("distribute_trade_msg_model", _distributeTradeMsgModel)
	return nil
}

// GetDistributeTradeMsgModel DistributeTradeMsgModel Getter
func (r AlibabaAlicomVtDistributeSendcodeAPIRequest) GetDistributeTradeMsgModel() *DistributeTradeMsgModel {
	return r._distributeTradeMsgModel
}

var poolAlibabaAlicomVtDistributeSendcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomVtDistributeSendcodeRequest()
	},
}

// GetAlibabaAlicomVtDistributeSendcodeRequest 从 sync.Pool 获取 AlibabaAlicomVtDistributeSendcodeAPIRequest
func GetAlibabaAlicomVtDistributeSendcodeAPIRequest() *AlibabaAlicomVtDistributeSendcodeAPIRequest {
	return poolAlibabaAlicomVtDistributeSendcodeAPIRequest.Get().(*AlibabaAlicomVtDistributeSendcodeAPIRequest)
}

// ReleaseAlibabaAlicomVtDistributeSendcodeAPIRequest 将 AlibabaAlicomVtDistributeSendcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomVtDistributeSendcodeAPIRequest(v *AlibabaAlicomVtDistributeSendcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlicomVtDistributeSendcodeAPIRequest.Put(v)
}
