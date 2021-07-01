package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalStoresanditemsAPIRequest
查询商户门店，商品列表 API请求
alibaba.alihealth.reserve.dental.storesanditems

查询商户门店，商品列表 */
type AlibabaAlihealthReserveDentalStoresanditemsAPIRequest struct {
	model.Params
	// 页码，每页100个门店，超过100个门店分页请求
	_pageNo int64
}

// NewAlibabaAlihealthReserveDentalStoresanditemsRequest 初始化AlibabaAlihealthReserveDentalStoresanditemsAPIRequest对象
func NewAlibabaAlihealthReserveDentalStoresanditemsRequest() *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest {
	return &AlibabaAlihealthReserveDentalStoresanditemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.storesanditems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageNo Setter
// 页码，每页100个门店，超过100个门店分页请求
func (r *AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaAlihealthReserveDentalStoresanditemsAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
