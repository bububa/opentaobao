package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcitydistributiongetAPIRequest 酒店城市数据获取接口-分销场景使用 API请求
// taobao.xhotel.city.distribution.get
//
// 引流API，对外提供酒店城市数据
type TaobaoxhotelcitydistributiongetAPIRequest struct {
	model.Params
	// 分页读取的开始下标,从0开始
	_start int64
	// 分页读取的城市个数，最小值为1，最大值为200
	_count int64
}

// NewTaobaoxhotelcitydistributiongetRequest 初始化TaobaoxhotelcitydistributiongetAPIRequest对象
func NewTaobaoxhotelcitydistributiongetRequest() *TaobaoxhotelcitydistributiongetAPIRequest {
	return &TaobaoxhotelcitydistributiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcitydistributiongetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.city.distribution.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcitydistributiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcitydistributiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 分页读取的开始下标,从0开始
func (r *TaobaoxhotelcitydistributiongetAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TaobaoxhotelcitydistributiongetAPIRequest) GetStart() int64 {
	return r._start
}

// SetCount is Count Setter
// 分页读取的城市个数，最小值为1，最大值为200
func (r *TaobaoxhotelcitydistributiongetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoxhotelcitydistributiongetAPIRequest) GetCount() int64 {
	return r._count
}
