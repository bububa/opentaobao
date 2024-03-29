package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcApContractsignedGetAPIRequest 用户是否签署支付宝代扣协议 API请求
// taobao.oc.ap.contractsigned.get
//
// 用户是否签署支付宝代扣协议
type TaobaoOcApContractsignedGetAPIRequest struct {
	model.Params
}

// NewTaobaoOcApContractsignedGetRequest 初始化TaobaoOcApContractsignedGetAPIRequest对象
func NewTaobaoOcApContractsignedGetRequest() *TaobaoOcApContractsignedGetAPIRequest {
	return &TaobaoOcApContractsignedGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOcApContractsignedGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcApContractsignedGetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.ap.contractsigned.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcApContractsignedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOcApContractsignedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoOcApContractsignedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOcApContractsignedGetRequest()
	},
}

// GetTaobaoOcApContractsignedGetRequest 从 sync.Pool 获取 TaobaoOcApContractsignedGetAPIRequest
func GetTaobaoOcApContractsignedGetAPIRequest() *TaobaoOcApContractsignedGetAPIRequest {
	return poolTaobaoOcApContractsignedGetAPIRequest.Get().(*TaobaoOcApContractsignedGetAPIRequest)
}

// ReleaseTaobaoOcApContractsignedGetAPIRequest 将 TaobaoOcApContractsignedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOcApContractsignedGetAPIRequest(v *TaobaoOcApContractsignedGetAPIRequest) {
	v.Reset()
	poolTaobaoOcApContractsignedGetAPIRequest.Put(v)
}
