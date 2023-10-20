package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalStoresanditemsAPIRequest 查询商户门店，商品列表 API请求
// alibaba.alihealth.reserve.dental.storesanditems
//
// 查询商户门店，商品列表
type AlibabaAlihealthReserveDentalStoresanditemsAPIRequest struct {
	model.Params
	// 页码，每页100个门店，超过100个门店分页请求
	_pageNo int64
}

// NewAlibabaAlihealthReserveDentalStoresanditemsRequest 初始化AlibabaAlihealthReserveDentalStoresanditemsAPIRequest对象
func NewAlibabaAlihealthReserveDentalStoresanditemsRequest() *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest {
	return &AlibabaAlihealthReserveDentalStoresanditemsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) Reset() {
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.storesanditems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 页码，每页100个门店，超过100个门店分页请求
func (r *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolAlibabaAlihealthReserveDentalStoresanditemsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthReserveDentalStoresanditemsRequest()
	},
}

// GetAlibabaAlihealthReserveDentalStoresanditemsRequest 从 sync.Pool 获取 AlibabaAlihealthReserveDentalStoresanditemsAPIRequest
func GetAlibabaAlihealthReserveDentalStoresanditemsAPIRequest() *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest {
	return poolAlibabaAlihealthReserveDentalStoresanditemsAPIRequest.Get().(*AlibabaAlihealthReserveDentalStoresanditemsAPIRequest)
}

// ReleaseAlibabaAlihealthReserveDentalStoresanditemsAPIRequest 将 AlibabaAlihealthReserveDentalStoresanditemsAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthReserveDentalStoresanditemsAPIRequest(v *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthReserveDentalStoresanditemsAPIRequest.Put(v)
}
