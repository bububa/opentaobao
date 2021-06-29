package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】改签通知单列表 APIRequest
taobao.alitrip.supplier.modify.list

提供供应商查询改签通知单列表
*/
type TaobaoAlitripSupplierModifyListRequest struct {
    model.Params

    // 页码
    currentPage   int64 

    // 乘客出发时间查询结束日期
    depEnd   string 

    // 乘客出发时间查询开始日期
    depStart   string 

    // 申请单创建时间查询结束日期
    gmtCreateEnd   string 

    // 申请单创建时间查询开始日期
    gmtCreateStart   string 

    // 每页记录数
    pageSize   int64 

    // 1：改签申请列表，2：改签已支付列表，3：改签成功列表
    status   int64 

}

func NewTaobaoAlitripSupplierModifyListRequest() *TaobaoAlitripSupplierModifyListRequest{
    return &TaobaoAlitripSupplierModifyListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSupplierModifyListRequest) GetApiMethodName() string {
    return "taobao.alitrip.supplier.modify.list"
}

func (r TaobaoAlitripSupplierModifyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSupplierModifyListRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoAlitripSupplierModifyListRequest) SetDepEnd(depEnd string) error {
    r.depEnd = depEnd
    r.Set("dep_end", depEnd)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetDepEnd() string {
    return r.depEnd
}

func (r *TaobaoAlitripSupplierModifyListRequest) SetDepStart(depStart string) error {
    r.depStart = depStart
    r.Set("dep_start", depStart)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetDepStart() string {
    return r.depStart
}

func (r *TaobaoAlitripSupplierModifyListRequest) SetGmtCreateEnd(gmtCreateEnd string) error {
    r.gmtCreateEnd = gmtCreateEnd
    r.Set("gmt_create_end", gmtCreateEnd)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetGmtCreateEnd() string {
    return r.gmtCreateEnd
}

func (r *TaobaoAlitripSupplierModifyListRequest) SetGmtCreateStart(gmtCreateStart string) error {
    r.gmtCreateStart = gmtCreateStart
    r.Set("gmt_create_start", gmtCreateStart)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetGmtCreateStart() string {
    return r.gmtCreateStart
}

func (r *TaobaoAlitripSupplierModifyListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoAlitripSupplierModifyListRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoAlitripSupplierModifyListRequest) GetStatus() int64 {
    return r.status
}

