package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单查询接口 APIRequest
tmall.servicecenter.workcard.query

工单查询接口
*/
type TmallServicecenterWorkcardQueryRequest struct {
    model.Params

    // 门店/网点id
    serviceStoreId   int64 

    // 核销码
    identifyCode   string 

    // 工单id
    id   int64 

    // 工单创建开始时间
    gmtCreateStart   string 

    // 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
    gmtCreateEnd   string 

    // 淘宝交易订单号。主订单或子订单均可
    bizOrderId   int64 

    // 当前页数
    currentPage   int64 

    // 每页大小
    pageSize   int64 

}

func NewTmallServicecenterWorkcardQueryRequest() *TmallServicecenterWorkcardQueryRequest{
    return &TmallServicecenterWorkcardQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.query"
}

func (r TmallServicecenterWorkcardQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardQueryRequest) SetServiceStoreId(serviceStoreId int64) error {
    r.serviceStoreId = serviceStoreId
    r.Set("service_store_id", serviceStoreId)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetServiceStoreId() int64 {
    return r.serviceStoreId
}

func (r *TmallServicecenterWorkcardQueryRequest) SetIdentifyCode(identifyCode string) error {
    r.identifyCode = identifyCode
    r.Set("identify_code", identifyCode)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetIdentifyCode() string {
    return r.identifyCode
}

func (r *TmallServicecenterWorkcardQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetId() int64 {
    return r.id
}

func (r *TmallServicecenterWorkcardQueryRequest) SetGmtCreateStart(gmtCreateStart string) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetGmtCreateStart() string {
    return r.gmtCreateStart
}

func (r *TmallServicecenterWorkcardQueryRequest) SetGmtCreateEnd(gmtCreateEnd string) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetGmtCreateEnd() string {
    return r.gmtCreateEnd
}

func (r *TmallServicecenterWorkcardQueryRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *TmallServicecenterWorkcardQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TmallServicecenterWorkcardQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallServicecenterWorkcardQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

