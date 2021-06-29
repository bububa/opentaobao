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
    cityCode   int64
    // 分页参数：当前页数，从1开始计数。<br/>默认值：1
    currentPage   int64
    // 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
    pageSize   int64
    // pid
    pid   string
    // 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
    shid   int64
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
func (r *TaobaoXhotelInfoListGetRequest) SetCityCode(cityCode int64) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r TaobaoXhotelInfoListGetRequest) GetCityCode() int64 {
    return r.cityCode
}
// CurrentPage Setter
// 分页参数：当前页数，从1开始计数。<br/>默认值：1
func (r *TaobaoXhotelInfoListGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoXhotelInfoListGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// PageSize Setter
// 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
func (r *TaobaoXhotelInfoListGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoXhotelInfoListGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// Pid Setter
// pid
func (r *TaobaoXhotelInfoListGetRequest) SetPid(pid string) error {
    r.pid = pid
    r.Set("pid", pid)
    return nil
}

// Pid Getter
func (r TaobaoXhotelInfoListGetRequest) GetPid() string {
    return r.pid
}
// Shid Setter
// 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
func (r *TaobaoXhotelInfoListGetRequest) SetShid(shid int64) error {
    r.shid = shid
    r.Set("shid", shid)
    return nil
}

// Shid Getter
func (r TaobaoXhotelInfoListGetRequest) GetShid() int64 {
    return r.shid
}
