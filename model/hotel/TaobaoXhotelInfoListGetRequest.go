package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店详细信息查询 API请求
taobao.xhotel.info.list.get

获取酒店详情信息
*/
type TaobaoXhotelInfoListGetRequest struct {
    model.Params
    // 城市code
    _cityCode   int64
    // 分页参数：当前页数，从1开始计数。<br/>默认值：1
    _currentPage   int64
    // 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
    _pageSize   int64
    // pid
    _pid   string
    // 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
    _shid   int64
}

// 初始化TaobaoXhotelInfoListGetRequest对象
func NewTaobaoXhotelInfoListGetRequest() *TaobaoXhotelInfoListGetRequest{
    return &TaobaoXhotelInfoListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelInfoListGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.info.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelInfoListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CityCode Setter
// 城市code
func (r *TaobaoXhotelInfoListGetRequest) SetCityCode(_cityCode int64) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoXhotelInfoListGetRequest) GetCityCode() int64 {
    return r._cityCode
}
// CurrentPage Setter
// 分页参数：当前页数，从1开始计数。<br/>默认值：1
func (r *TaobaoXhotelInfoListGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoXhotelInfoListGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
func (r *TaobaoXhotelInfoListGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelInfoListGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// Pid Setter
// pid
func (r *TaobaoXhotelInfoListGetRequest) SetPid(_pid string) error {
    r._pid = _pid
    r.Set("pid", _pid)
    return nil
}

// Pid Getter
func (r TaobaoXhotelInfoListGetRequest) GetPid() string {
    return r._pid
}
// Shid Setter
// 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
func (r *TaobaoXhotelInfoListGetRequest) SetShid(_shid int64) error {
    r._shid = _shid
    r.Set("shid", _shid)
    return nil
}

// Shid Getter
func (r TaobaoXhotelInfoListGetRequest) GetShid() int64 {
    return r._shid
}
