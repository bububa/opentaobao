package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest 阿信酒店分销-标准酒店id列表查询 API请求
// taobao.alitrip.travel.axin.hotel.shid.list.query
//
// 标准酒店id列表查询
type TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest struct {
	model.Params
	// 资源渠道
	_resourceChannel string
	// 页码
	_pageNo int64
	// 页大小
	_pageSize int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelShidListQueryRequest 初始化TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelShidListQueryRequest() *TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest {
	return &TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.shid.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
