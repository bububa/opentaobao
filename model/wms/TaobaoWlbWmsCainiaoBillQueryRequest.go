package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单据列表 APIRequest
taobao.wlb.wms.cainiao.bill.query

查询单据列表
*/
type TaobaoWlbWmsCainiaoBillQueryRequest struct {
    model.Params

    // 结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
    startModifiedTime   string 

    // 起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
    endModifiedTime   string 

    // 订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
    orderType   string 

    // 页码。（大于0的整数。默认为1）
    pageNo   int64 

    // 每页条数。（每页条数不超过50条。默认为50）
    pageSize   int64 

}

func NewTaobaoWlbWmsCainiaoBillQueryRequest() *TaobaoWlbWmsCainiaoBillQueryRequest{
    return &TaobaoWlbWmsCainiaoBillQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.cainiao.bill.query"
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetStartModifiedTime(startModifiedTime string) error {
    r.startModifiedTime = startModifiedTime
    r.Set("start_modified_time", startModifiedTime)
    return nil
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetStartModifiedTime() string {
    return r.startModifiedTime
}

func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetEndModifiedTime(endModifiedTime string) error {
    r.endModifiedTime = endModifiedTime
    r.Set("end_modified_time", endModifiedTime)
    return nil
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetEndModifiedTime() string {
    return r.endModifiedTime
}

func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetOrderType() string {
    return r.orderType
}

func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

