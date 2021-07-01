package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelCityGetAPIRequest
酒店城市数据获取接口 API请求
taobao.xhotel.city.get

引流API，对外提供酒店城市数据 */
type TaobaoXhotelCityGetAPIRequest struct {
	model.Params
	// 分页读取的开始下标,从0开始
	_start int64
	// 分页读取的城市个数，最小值为1，最大值为200
	_count int64
}

// NewTaobaoXhotelCityGetRequest 初始化TaobaoXhotelCityGetAPIRequest对象
func NewTaobaoXhotelCityGetRequest() *TaobaoXhotelCityGetAPIRequest {
	return &TaobaoXhotelCityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Start Setter
// 分页读取的开始下标,从0开始
func (r *TaobaoXhotelCityGetAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// Get Start Getter
func (r TaobaoXhotelCityGetAPIRequest) GetStart() int64 {
	return r._start
}

// Set is Count Setter
// 分页读取的城市个数，最小值为1，最大值为200
func (r *TaobaoXhotelCityGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// Get Count Getter
func (r TaobaoXhotelCityGetAPIRequest) GetCount() int64 {
	return r._count
}
