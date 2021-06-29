package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过物流订单编号查询物流信息 API请求
taobao.wlb.tmsorder.query

通过物流订单编号分页查询物流信息
*/
type TaobaoWlbTmsorderQueryRequest struct {
    model.Params
    // 物流订单编号
    orderCode   string
    // 当前页
    pageNo   int64
    // 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    pageSize   int64
}

// 初始化TaobaoWlbTmsorderQueryRequest对象
func NewTaobaoWlbTmsorderQueryRequest() *TaobaoWlbTmsorderQueryRequest{
    return &TaobaoWlbTmsorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbTmsorderQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.tmsorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbTmsorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 物流订单编号
func (r *TaobaoWlbTmsorderQueryRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbTmsorderQueryRequest) GetOrderCode() string {
    return r.orderCode
}
// PageNo Setter
// 当前页
func (r *TaobaoWlbTmsorderQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbTmsorderQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
func (r *TaobaoWlbTmsorderQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbTmsorderQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
