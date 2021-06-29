package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单据列表 API请求
taobao.wlb.wms.cainiao.bill.query

查询单据列表
*/
type TaobaoWlbWmsCainiaoBillQueryRequest struct {
    model.Params
    // 结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
    _startModifiedTime   string
    // 起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
    _endModifiedTime   string
    // 订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
    _orderType   string
    // 页码。（大于0的整数。默认为1）
    _pageNo   int64
    // 每页条数。（每页条数不超过50条。默认为50）
    _pageSize   int64
}

// 初始化TaobaoWlbWmsCainiaoBillQueryRequest对象
func NewTaobaoWlbWmsCainiaoBillQueryRequest() *TaobaoWlbWmsCainiaoBillQueryRequest{
    return &TaobaoWlbWmsCainiaoBillQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.cainiao.bill.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartModifiedTime Setter
// 结束时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetStartModifiedTime(_startModifiedTime string) error {
    r._startModifiedTime = _startModifiedTime
    r.Set("start_modified_time", _startModifiedTime)
    return nil
}

// StartModifiedTime Getter
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetStartModifiedTime() string {
    return r._startModifiedTime
}
// EndModifiedTime Setter
// 起始时间，此字段检索订单最后修改时间， 格式 yyyy-MM-dd HH:mm:ss。
func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetEndModifiedTime(_endModifiedTime string) error {
    r._endModifiedTime = _endModifiedTime
    r.Set("end_modified_time", _endModifiedTime)
    return nil
}

// EndModifiedTime Getter
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetEndModifiedTime() string {
    return r._endModifiedTime
}
// OrderType Setter
// 订单类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetOrderType(_orderType string) error {
    r._orderType = _orderType
    r.Set("order_type", _orderType)
    return nil
}

// OrderType Getter
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetOrderType() string {
    return r._orderType
}
// PageNo Setter
// 页码。（大于0的整数。默认为1）
func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数。（每页条数不超过50条。默认为50）
func (r *TaobaoWlbWmsCainiaoBillQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbWmsCainiaoBillQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
