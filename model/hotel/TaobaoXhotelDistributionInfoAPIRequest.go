package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldistributioninfoAPIRequest 飞猪分销通用酒店标准信息接口 API请求
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
type TaobaoxhoteldistributioninfoAPIRequest struct {
	model.Params
	// pid
	_pid string
	// 城市code
	_cityCode int64
	// 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
	_shid int64
	// 分页参数：当前页数，从1开始计数。默认值：1
	_currentPage int64
	// 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
	_pageSize int64
}

// NewTaobaoxhoteldistributioninfoRequest 初始化TaobaoxhoteldistributioninfoAPIRequest对象
func NewTaobaoxhoteldistributioninfoRequest() *TaobaoxhoteldistributioninfoAPIRequest {
	return &TaobaoxhoteldistributioninfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhoteldistributioninfoAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.distribution.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhoteldistributioninfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhoteldistributioninfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPid is Pid Setter
// pid
func (r *TaobaoxhoteldistributioninfoAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaoxhoteldistributioninfoAPIRequest) GetPid() string {
	return r._pid
}

// SetCityCode is CityCode Setter
// 城市code
func (r *TaobaoxhoteldistributioninfoAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoxhoteldistributioninfoAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetShid is Shid Setter
// 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
func (r *TaobaoxhoteldistributioninfoAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoxhoteldistributioninfoAPIRequest) GetShid() int64 {
	return r._shid
}

// SetCurrentPage is CurrentPage Setter
// 分页参数：当前页数，从1开始计数。默认值：1
func (r *TaobaoxhoteldistributioninfoAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoxhoteldistributioninfoAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
func (r *TaobaoxhoteldistributioninfoAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoxhoteldistributioninfoAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
