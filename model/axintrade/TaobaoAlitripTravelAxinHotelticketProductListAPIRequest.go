package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketproductlistAPIRequest 阿信酒景套餐产品列表查询 API请求
// taobao.alitrip.travel.axin.hotelticket.product.list
//
// 阿信酒景套餐产品列表查询
type TaobaoalitriptravelaxinhotelticketproductlistAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 当前页数
	_pageNo int64
	// 分页大小
	_pageSize int64
}

// NewTaobaoalitriptravelaxinhotelticketproductlistRequest 初始化TaobaoalitriptravelaxinhotelticketproductlistAPIRequest对象
func NewTaobaoalitriptravelaxinhotelticketproductlistRequest() *TaobaoalitriptravelaxinhotelticketproductlistAPIRequest {
	return &TaobaoalitriptravelaxinhotelticketproductlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPageNo is PageNo Setter
// 当前页数
func (r *TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoalitriptravelaxinhotelticketproductlistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
