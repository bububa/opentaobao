package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单据序列号信息 APIRequest
taobao.wlb.wms.sn.info.query

查询仓库作业的各类单据记录的Sn信息
*/
type TaobaoWlbWmsSnInfoQueryRequest struct {
    model.Params

    // 订单编码
    orderCode   string 

    // 订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
    orderCodeType   int64 

    // 页数，默认每页50条
    pageIndex   int64 

}

func NewTaobaoWlbWmsSnInfoQueryRequest() *TaobaoWlbWmsSnInfoQueryRequest{
    return &TaobaoWlbWmsSnInfoQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsSnInfoQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.sn.info.query"
}

func (r TaobaoWlbWmsSnInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsSnInfoQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsSnInfoQueryRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbWmsSnInfoQueryRequest) SetOrderCodeType(orderCodeType int64) error {
    r.orderCodeType = orderCodeType
    r.Set("order_code_type", orderCodeType)
    return nil
}

func (r TaobaoWlbWmsSnInfoQueryRequest) GetOrderCodeType() int64 {
    return r.orderCodeType
}

func (r *TaobaoWlbWmsSnInfoQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoWlbWmsSnInfoQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}

