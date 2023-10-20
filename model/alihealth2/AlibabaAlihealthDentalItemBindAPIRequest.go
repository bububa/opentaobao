package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemBindAPIRequest ISV绑定外部门店id和外部商品id API请求
// alibaba.alihealth.dental.item.bind
//
// ISV绑定外部门店id和外部商品id
type AlibabaAlihealthDentalItemBindAPIRequest struct {
	model.Params
	// bind_list
	_bindList []StoreItemRelRequest
	// 类型 1 天猫门店 2 支付宝门店
	_type int64
}

// NewAlibabaAlihealthDentalItemBindRequest 初始化AlibabaAlihealthDentalItemBindAPIRequest对象
func NewAlibabaAlihealthDentalItemBindRequest() *AlibabaAlihealthDentalItemBindAPIRequest {
	return &AlibabaAlihealthDentalItemBindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDentalItemBindAPIRequest) Reset() {
	r._bindList = r._bindList[:0]
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.item.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindList is BindList Setter
// bind_list
func (r *AlibabaAlihealthDentalItemBindAPIRequest) SetBindList(_bindList []StoreItemRelRequest) error {
	r._bindList = _bindList
	r.Set("bind_list", _bindList)
	return nil
}

// GetBindList BindList Getter
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetBindList() []StoreItemRelRequest {
	return r._bindList
}

// SetType is Type Setter
// 类型 1 天猫门店 2 支付宝门店
func (r *AlibabaAlihealthDentalItemBindAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetType() int64 {
	return r._type
}

var poolAlibabaAlihealthDentalItemBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDentalItemBindRequest()
	},
}

// GetAlibabaAlihealthDentalItemBindRequest 从 sync.Pool 获取 AlibabaAlihealthDentalItemBindAPIRequest
func GetAlibabaAlihealthDentalItemBindAPIRequest() *AlibabaAlihealthDentalItemBindAPIRequest {
	return poolAlibabaAlihealthDentalItemBindAPIRequest.Get().(*AlibabaAlihealthDentalItemBindAPIRequest)
}

// ReleaseAlibabaAlihealthDentalItemBindAPIRequest 将 AlibabaAlihealthDentalItemBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDentalItemBindAPIRequest(v *AlibabaAlihealthDentalItemBindAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDentalItemBindAPIRequest.Put(v)
}
