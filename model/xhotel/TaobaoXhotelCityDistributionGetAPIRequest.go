package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCityDistributionGetAPIRequest 酒店城市数据获取接口-分销场景使用 API请求
// taobao.xhotel.city.distribution.get
//
// 引流API，对外提供酒店城市数据
type TaobaoXhotelCityDistributionGetAPIRequest struct {
	model.Params
	// 分页读取的开始下标,从0开始
	_start int64
	// 分页读取的城市个数，最小值为1，最大值为200
	_count int64
}

// NewTaobaoXhotelCityDistributionGetRequest 初始化TaobaoXhotelCityDistributionGetAPIRequest对象
func NewTaobaoXhotelCityDistributionGetRequest() *TaobaoXhotelCityDistributionGetAPIRequest {
	return &TaobaoXhotelCityDistributionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityDistributionGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.city.distribution.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityDistributionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelCityDistributionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 分页读取的开始下标,从0开始
func (r *TaobaoXhotelCityDistributionGetAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaoXhotelCityDistributionGetAPIRequest) GetStart() int64 {
	return r._start
}

// SetCount is Count Setter
// 分页读取的城市个数，最小值为1，最大值为200
func (r *TaobaoXhotelCityDistributionGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoXhotelCityDistributionGetAPIRequest) GetCount() int64 {
	return r._count
}
