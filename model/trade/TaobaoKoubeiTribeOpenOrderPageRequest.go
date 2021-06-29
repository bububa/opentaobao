package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体订单列表信息查询 API请求
taobao.koubei.tribe.open.order.page

查询口碑商圈用户的订单列表信息
*/
type TaobaoKoubeiTribeOpenOrderPageRequest struct {
    model.Params
    // 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
    _orderStatus   string
    // 每页大小
    _pageSize   int64
    // 起始页
    _pageNo   int64
    // 数据集Id
    _dataSetId   string
    // 用户openId
    _openId   string
}

// 初始化TaobaoKoubeiTribeOpenOrderPageRequest对象
func NewTaobaoKoubeiTribeOpenOrderPageRequest() *TaobaoKoubeiTribeOpenOrderPageRequest{
    return &TaobaoKoubeiTribeOpenOrderPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetApiMethodName() string {
    return "taobao.koubei.tribe.open.order.page"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderStatus Setter
// 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetOrderStatus(_orderStatus string) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetOrderStatus() string {
    return r._orderStatus
}
// PageSize Setter
// 每页大小
func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 起始页
func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetPageNo() int64 {
    return r._pageNo
}
// DataSetId Setter
// 数据集Id
func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetDataSetId() string {
    return r._dataSetId
}
// OpenId Setter
// 用户openId
func (r *TaobaoKoubeiTribeOpenOrderPageRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r TaobaoKoubeiTribeOpenOrderPageRequest) GetOpenId() string {
    return r._openId
}
