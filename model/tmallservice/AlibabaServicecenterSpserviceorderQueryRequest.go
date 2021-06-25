package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单查询 APIRequest
alibaba.servicecenter.spserviceorder.query

服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
*/
type AlibabaServicecenterSpserviceorderQueryRequest struct {
    model.Params

    // 状态码，可传多个
    statusCodes   string 

    // 查询第几页，默认1
    currentPage   int64 

    // 每页大小，默认50，最大50
    pageSize   int64 

    // 服务单修改时间(时间段15分钟以内)(包含时分秒)
    gmtModifiedEnd   string 

    // 服务单修改时间(包含时分秒)
    gmtModifiedStart   string 

    // 实物主订单id(消费者在淘宝订单里看到的订单号)
    masterParentBizOrderId   int64 

    // 服务单id
    spServiceOrderId   int64 

}

func NewAlibabaServicecenterSpserviceorderQueryRequest() *AlibabaServicecenterSpserviceorderQueryRequest{
    return &AlibabaServicecenterSpserviceorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.spserviceorder.query"
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetStatusCodes(statusCodes string) error {
    r.statusCodes = statusCodes
    r.Set("status_codes", statusCodes)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetStatusCodes() string {
    return r.statusCodes
}

func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetGmtModifiedEnd(gmtModifiedEnd string) error {
    r.gmtModifiedEnd = gmtModifiedEnd
    r.Set("gmt_modified_end", gmtModifiedEnd)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetGmtModifiedEnd() string {
    return r.gmtModifiedEnd
}

func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetGmtModifiedStart(gmtModifiedStart string) error {
    r.gmtModifiedStart = gmtModifiedStart
    r.Set("gmt_modified_start", gmtModifiedStart)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetGmtModifiedStart() string {
    return r.gmtModifiedStart
}

func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetMasterParentBizOrderId(masterParentBizOrderId int64) error {
    r.masterParentBizOrderId = masterParentBizOrderId
    r.Set("master_parent_biz_order_id", masterParentBizOrderId)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetMasterParentBizOrderId() int64 {
    return r.masterParentBizOrderId
}

func (r *AlibabaServicecenterSpserviceorderQueryRequest) SetSpServiceOrderId(spServiceOrderId int64) error {
    r.spServiceOrderId = spServiceOrderId
    r.Set("sp_service_order_id", spServiceOrderId)
    return nil
}

func (r AlibabaServicecenterSpserviceorderQueryRequest) GetSpServiceOrderId() int64 {
    return r.spServiceOrderId
}

