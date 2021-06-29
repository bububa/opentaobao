package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID查询所有库存变更记录 API请求
taobao.wlb.inventorylog.query

通过商品ID等几个条件来分页查询库存变更记录
*/
type TaobaoWlbInventorylogQueryRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 仓库编码
    _storeCode   string
    // 单号
    _orderCode   string
    // 起始修改时间,大于等于该时间
    _gmtStart   string
    // 结束修改时间,小于等于该时间
    _gmtEnd   string
    // 当前页
    _pageNo   int64
    // 分页记录个数
    _pageSize   int64
    // 可指定授权的用户来查询
    _opUserId   int64
    // 库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
    _opType   string
}

// 初始化TaobaoWlbInventorylogQueryRequest对象
func NewTaobaoWlbInventorylogQueryRequest() *TaobaoWlbInventorylogQueryRequest{
    return &TaobaoWlbInventorylogQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbInventorylogQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.inventorylog.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbInventorylogQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoWlbInventorylogQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbInventorylogQueryRequest) GetItemId() int64 {
    return r._itemId
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbInventorylogQueryRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbInventorylogQueryRequest) GetStoreCode() string {
    return r._storeCode
}
// OrderCode Setter
// 单号
func (r *TaobaoWlbInventorylogQueryRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbInventorylogQueryRequest) GetOrderCode() string {
    return r._orderCode
}
// GmtStart Setter
// 起始修改时间,大于等于该时间
func (r *TaobaoWlbInventorylogQueryRequest) SetGmtStart(_gmtStart string) error {
    r._gmtStart = _gmtStart
    r.Set("gmt_start", _gmtStart)
    return nil
}

// GmtStart Getter
func (r TaobaoWlbInventorylogQueryRequest) GetGmtStart() string {
    return r._gmtStart
}
// GmtEnd Setter
// 结束修改时间,小于等于该时间
func (r *TaobaoWlbInventorylogQueryRequest) SetGmtEnd(_gmtEnd string) error {
    r._gmtEnd = _gmtEnd
    r.Set("gmt_end", _gmtEnd)
    return nil
}

// GmtEnd Getter
func (r TaobaoWlbInventorylogQueryRequest) GetGmtEnd() string {
    return r._gmtEnd
}
// PageNo Setter
// 当前页
func (r *TaobaoWlbInventorylogQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbInventorylogQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页记录个数
func (r *TaobaoWlbInventorylogQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbInventorylogQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// OpUserId Setter
// 可指定授权的用户来查询
func (r *TaobaoWlbInventorylogQueryRequest) SetOpUserId(_opUserId int64) error {
    r._opUserId = _opUserId
    r.Set("op_user_id", _opUserId)
    return nil
}

// OpUserId Getter
func (r TaobaoWlbInventorylogQueryRequest) GetOpUserId() int64 {
    return r._opUserId
}
// OpType Setter
// 库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
func (r *TaobaoWlbInventorylogQueryRequest) SetOpType(_opType string) error {
    r._opType = _opType
    r.Set("op_type", _opType)
    return nil
}

// OpType Getter
func (r TaobaoWlbInventorylogQueryRequest) GetOpType() string {
    return r._opType
}
