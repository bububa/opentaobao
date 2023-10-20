package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalBindshopanditemAPIRequest 绑定门店信息，商品信息 API请求
// alibaba.alihealth.reserve.dental.bindshopanditem
//
// 绑定门店信息，商品信息
type AlibabaAlihealthReserveDentalBindshopanditemAPIRequest struct {
	model.Params
	// bind_list
	_bindList []BindDto
}

// NewAlibabaAlihealthReserveDentalBindshopanditemRequest 初始化AlibabaAlihealthReserveDentalBindshopanditemAPIRequest对象
func NewAlibabaAlihealthReserveDentalBindshopanditemRequest() *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest {
	return &AlibabaAlihealthReserveDentalBindshopanditemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) Reset() {
	r._bindList = r._bindList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.bindshopanditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindList is BindList Setter
// bind_list
func (r *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) SetBindList(_bindList []BindDto) error {
	r._bindList = _bindList
	r.Set("bind_list", _bindList)
	return nil
}

// GetBindList BindList Getter
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetBindList() []BindDto {
	return r._bindList
}

var poolAlibabaAlihealthReserveDentalBindshopanditemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthReserveDentalBindshopanditemRequest()
	},
}

// GetAlibabaAlihealthReserveDentalBindshopanditemRequest 从 sync.Pool 获取 AlibabaAlihealthReserveDentalBindshopanditemAPIRequest
func GetAlibabaAlihealthReserveDentalBindshopanditemAPIRequest() *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest {
	return poolAlibabaAlihealthReserveDentalBindshopanditemAPIRequest.Get().(*AlibabaAlihealthReserveDentalBindshopanditemAPIRequest)
}

// ReleaseAlibabaAlihealthReserveDentalBindshopanditemAPIRequest 将 AlibabaAlihealthReserveDentalBindshopanditemAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthReserveDentalBindshopanditemAPIRequest(v *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthReserveDentalBindshopanditemAPIRequest.Put(v)
}
