package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelcitygetAPIRequest 城市列表信息查询-阿信 API请求
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
type TaobaoalitriptravelaxinhotelcitygetAPIRequest struct {
	model.Params
	// 分销商id，阿信分配
	_distributorTid int64
}

// NewTaobaoalitriptravelaxinhotelcitygetRequest 初始化TaobaoalitriptravelaxinhotelcitygetAPIRequest对象
func NewTaobaoalitriptravelaxinhotelcitygetRequest() *TaobaoalitriptravelaxinhotelcitygetAPIRequest {
	return &TaobaoalitriptravelaxinhotelcitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelcitygetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelcitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelcitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id，阿信分配
func (r *TaobaoalitriptravelaxinhotelcitygetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelcitygetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
