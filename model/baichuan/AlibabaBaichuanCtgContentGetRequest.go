package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川内容平台内容获取 API请求
alibaba.baichuan.ctg.content.get

百川内容平台内容获取
*/
type AlibabaBaichuanCtgContentGetRequest struct {
    model.Params
    // 投放位置
    _deliveryId   string
    // 分页大小
    _pageSize   int64
    // 当前页
    _currentPage   int64
    // 资源位
    _resId   string
    // 日期
    _date   string
}

// 初始化AlibabaBaichuanCtgContentGetRequest对象
func NewAlibabaBaichuanCtgContentGetRequest() *AlibabaBaichuanCtgContentGetRequest{
    return &AlibabaBaichuanCtgContentGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgContentGetRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.content.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgContentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryId Setter
// 投放位置
func (r *AlibabaBaichuanCtgContentGetRequest) SetDeliveryId(_deliveryId string) error {
    r._deliveryId = _deliveryId
    r.Set("delivery_id", _deliveryId)
    return nil
}

// DeliveryId Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetDeliveryId() string {
    return r._deliveryId
}
// PageSize Setter
// 分页大小
func (r *AlibabaBaichuanCtgContentGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 当前页
func (r *AlibabaBaichuanCtgContentGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// ResId Setter
// 资源位
func (r *AlibabaBaichuanCtgContentGetRequest) SetResId(_resId string) error {
    r._resId = _resId
    r.Set("res_id", _resId)
    return nil
}

// ResId Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetResId() string {
    return r._resId
}
// Date Setter
// 日期
func (r *AlibabaBaichuanCtgContentGetRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetDate() string {
    return r._date
}
