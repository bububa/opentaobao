package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsResourceGetAPIRequest 获取所有服务列表 API请求
// taobao.wlb.imports.resource.get
//
// 一般进口TOP接口，获取所有服务列表。
type TaobaoWlbImportsResourceGetAPIRequest struct {
	model.Params
	// 卖家发货地区域ID
	_fromId int64
	// 买家收货地信息
	_toAddress *ReciverAddressDo
}

// NewTaobaoWlbImportsResourceGetRequest 初始化TaobaoWlbImportsResourceGetAPIRequest对象
func NewTaobaoWlbImportsResourceGetRequest() *TaobaoWlbImportsResourceGetAPIRequest {
	return &TaobaoWlbImportsResourceGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbImportsResourceGetAPIRequest) Reset() {
	r._fromId = 0
	r._toAddress = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsResourceGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.resource.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsResourceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbImportsResourceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFromId is FromId Setter
// 卖家发货地区域ID
func (r *TaobaoWlbImportsResourceGetAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r TaobaoWlbImportsResourceGetAPIRequest) GetFromId() int64 {
	return r._fromId
}

// SetToAddress is ToAddress Setter
// 买家收货地信息
func (r *TaobaoWlbImportsResourceGetAPIRequest) SetToAddress(_toAddress *ReciverAddressDo) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r TaobaoWlbImportsResourceGetAPIRequest) GetToAddress() *ReciverAddressDo {
	return r._toAddress
}

var poolTaobaoWlbImportsResourceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbImportsResourceGetRequest()
	},
}

// GetTaobaoWlbImportsResourceGetRequest 从 sync.Pool 获取 TaobaoWlbImportsResourceGetAPIRequest
func GetTaobaoWlbImportsResourceGetAPIRequest() *TaobaoWlbImportsResourceGetAPIRequest {
	return poolTaobaoWlbImportsResourceGetAPIRequest.Get().(*TaobaoWlbImportsResourceGetAPIRequest)
}

// ReleaseTaobaoWlbImportsResourceGetAPIRequest 将 TaobaoWlbImportsResourceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbImportsResourceGetAPIRequest(v *TaobaoWlbImportsResourceGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbImportsResourceGetAPIRequest.Put(v)
}
