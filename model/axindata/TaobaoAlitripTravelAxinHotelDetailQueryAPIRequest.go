package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhoteldetailqueryAPIRequest 阿信酒店分销-标准酒店详情查询 API请求
// taobao.alitrip.travel.axin.hotel.detail.query
//
// 标准酒店详情查询
type TaobaoalitriptravelaxinhoteldetailqueryAPIRequest struct {
	model.Params
	// 资源渠道
	_resourceChannel string
	// 标准酒店id
	_shid int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoalitriptravelaxinhoteldetailqueryRequest 初始化TaobaoalitriptravelaxinhoteldetailqueryAPIRequest对象
func NewTaobaoalitriptravelaxinhoteldetailqueryRequest() *TaobaoalitriptravelaxinhoteldetailqueryAPIRequest {
	return &TaobaoalitriptravelaxinhoteldetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhoteldetailqueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
