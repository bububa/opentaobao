package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlicomVtDistributeorderCreateAPIRequest) Reset() {
	r._distributeTradeMsgModel = nil
	r.Params.ToZero()
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

var poolAlibabaAlicomVtDistributeorderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlicomVtDistributeorderCreateRequest()
	},
}

// GetAlibabaAlicomVtDistributeorderCreateRequest 从 sync.Pool 获取 AlibabaAlicomVtDistributeorderCreateAPIRequest
func GetAlibabaAlicomVtDistributeorderCreateAPIRequest() *AlibabaAlicomVtDistributeorderCreateAPIRequest {
	return poolAlibabaAlicomVtDistributeorderCreateAPIRequest.Get().(*AlibabaAlicomVtDistributeorderCreateAPIRequest)
}

// ReleaseAlibabaAlicomVtDistributeorderCreateAPIRequest 将 AlibabaAlicomVtDistributeorderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlicomVtDistributeorderCreateAPIRequest(v *AlibabaAlicomVtDistributeorderCreateAPIRequest) {
	v.Reset()
	poolAlibabaAlicomVtDistributeorderCreateAPIRequest.Put(v)
}
