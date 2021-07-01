package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单据序列号信息 API请求
taobao.wlb.wms.sn.info.query

查询仓库作业的各类单据记录的Sn信息
*/
type TaobaoWlbWmsSnInfoQueryAPIRequest struct {
    model.Params
    // 订单编码
    _orderCode   string
    // 订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
    _orderCodeType   int64
    // 页数，默认每页50条
    _pageIndex   int64
}

// 初始化TaobaoWlbWmsSnInfoQueryAPIRequest对象
func NewTaobaoWlbWmsSnInfoQueryRequest() *TaobaoWlbWmsSnInfoQueryAPIRequest{
    return &TaobaoWlbWmsSnInfoQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsSnInfoQueryAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.sn.info.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsSnInfoQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 订单编码
func (r *TaobaoWlbWmsSnInfoQueryAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsSnInfoQueryAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// OrderCodeType Setter
// 订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
func (r *TaobaoWlbWmsSnInfoQueryAPIRequest) SetOrderCodeType(_orderCodeType int64) error {
    r._orderCodeType = _orderCodeType
    r.Set("order_code_type", _orderCodeType)
    return nil
}

// OrderCodeType Getter
func (r TaobaoWlbWmsSnInfoQueryAPIRequest) GetOrderCodeType() int64 {
    return r._orderCodeType
}
// PageIndex Setter
// 页数，默认每页50条
func (r *TaobaoWlbWmsSnInfoQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoWlbWmsSnInfoQueryAPIRequest) GetPageIndex() int64 {
    return r._pageIndex
}
