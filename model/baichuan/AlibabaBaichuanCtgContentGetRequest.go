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
    deliveryId   string
    // 分页大小
    pageSize   int64
    // 当前页
    currentPage   int64
    // 资源位
    resId   string
    // 日期
    date   string
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
func (r *AlibabaBaichuanCtgContentGetRequest) SetDeliveryId(deliveryId string) error {
    r.deliveryId = deliveryId
    r.Set("delivery_id", deliveryId)
    return nil
}

// DeliveryId Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetDeliveryId() string {
    return r.deliveryId
}
// PageSize Setter
// 分页大小
func (r *AlibabaBaichuanCtgContentGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// CurrentPage Setter
// 当前页
func (r *AlibabaBaichuanCtgContentGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// ResId Setter
// 资源位
func (r *AlibabaBaichuanCtgContentGetRequest) SetResId(resId string) error {
    r.resId = resId
    r.Set("res_id", resId)
    return nil
}

// ResId Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetResId() string {
    return r.resId
}
// Date Setter
// 日期
func (r *AlibabaBaichuanCtgContentGetRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r AlibabaBaichuanCtgContentGetRequest) GetDate() string {
    return r.date
}
