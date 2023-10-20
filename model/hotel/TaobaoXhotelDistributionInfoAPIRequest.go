package hotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelDistributionInfoAPIRequest 飞猪分销通用酒店标准信息接口 API请求
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
type TaobaoXhotelDistributionInfoAPIRequest struct {
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

// NewTaobaoXhotelDistributionInfoRequest 初始化TaobaoXhotelDistributionInfoAPIRequest对象
func NewTaobaoXhotelDistributionInfoRequest() *TaobaoXhotelDistributionInfoAPIRequest {
	return &TaobaoXhotelDistributionInfoAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelDistributionInfoAPIRequest) Reset() {
	r._pid = ""
	r._cityCode = 0
	r._shid = 0
	r._currentPage = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelDistributionInfoAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.distribution.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelDistributionInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelDistributionInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPid is Pid Setter
// pid
func (r *TaobaoXhotelDistributionInfoAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaoXhotelDistributionInfoAPIRequest) GetPid() string {
	return r._pid
}

// SetCityCode is CityCode Setter
// 城市code
func (r *TaobaoXhotelDistributionInfoAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoXhotelDistributionInfoAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetShid is Shid Setter
// 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
func (r *TaobaoXhotelDistributionInfoAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoXhotelDistributionInfoAPIRequest) GetShid() int64 {
	return r._shid
}

// SetCurrentPage is CurrentPage Setter
// 分页参数：当前页数，从1开始计数。默认值：1
func (r *TaobaoXhotelDistributionInfoAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoXhotelDistributionInfoAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
func (r *TaobaoXhotelDistributionInfoAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoXhotelDistributionInfoAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoXhotelDistributionInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelDistributionInfoRequest()
	},
}

// GetTaobaoXhotelDistributionInfoRequest 从 sync.Pool 获取 TaobaoXhotelDistributionInfoAPIRequest
func GetTaobaoXhotelDistributionInfoAPIRequest() *TaobaoXhotelDistributionInfoAPIRequest {
	return poolTaobaoXhotelDistributionInfoAPIRequest.Get().(*TaobaoXhotelDistributionInfoAPIRequest)
}

// ReleaseTaobaoXhotelDistributionInfoAPIRequest 将 TaobaoXhotelDistributionInfoAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelDistributionInfoAPIRequest(v *TaobaoXhotelDistributionInfoAPIRequest) {
	v.Reset()
	poolTaobaoXhotelDistributionInfoAPIRequest.Put(v)
}
