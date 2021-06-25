package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川内容平台内容获取 APIRequest
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

func NewAlibabaBaichuanCtgContentGetRequest() *AlibabaBaichuanCtgContentGetRequest{
    return &AlibabaBaichuanCtgContentGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanCtgContentGetRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.content.get"
}

func (r AlibabaBaichuanCtgContentGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBaichuanCtgContentGetRequest) SetDeliveryId(deliveryId string) error {
    r.deliveryId = deliveryId
    r.Set("delivery_id", deliveryId)
    return nil
}

func (r AlibabaBaichuanCtgContentGetRequest) GetDeliveryId() string {
    return r.deliveryId
}

func (r *AlibabaBaichuanCtgContentGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaBaichuanCtgContentGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaBaichuanCtgContentGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaBaichuanCtgContentGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaBaichuanCtgContentGetRequest) SetResId(resId string) error {
    r.resId = resId
    r.Set("res_id", resId)
    return nil
}

func (r AlibabaBaichuanCtgContentGetRequest) GetResId() string {
    return r.resId
}

func (r *AlibabaBaichuanCtgContentGetRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r AlibabaBaichuanCtgContentGetRequest) GetDate() string {
    return r.date
}

