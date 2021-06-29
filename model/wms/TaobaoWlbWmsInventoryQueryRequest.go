package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟商品库存查询 API请求
taobao.wlb.wms.inventory.query

支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
*/
type TaobaoWlbWmsInventoryQueryRequest struct {
    model.Params
    // 菜鸟商品ID
    _itemId   string
    // 仓库编码
    _storeCode   string
    // 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
    _inventoryType   int64
    // 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
    _type   int64
    // 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
    _batchCode   string
    // 生产日期，type=2时字段传值有效。
    _produceDate   string
    // 失效日期，type=2时字段传值有效。
    _dueDate   string
    // 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
    _channelCode   string
    // 分页的第几页
    _pageNo   int64
    // 每页多少条，最大50条
    _pageSize   int64
}

// 初始化TaobaoWlbWmsInventoryQueryRequest对象
func NewTaobaoWlbWmsInventoryQueryRequest() *TaobaoWlbWmsInventoryQueryRequest{
    return &TaobaoWlbWmsInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 菜鸟商品ID
func (r *TaobaoWlbWmsInventoryQueryRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetItemId() string {
    return r._itemId
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsInventoryQueryRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetStoreCode() string {
    return r._storeCode
}
// InventoryType Setter
// 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
func (r *TaobaoWlbWmsInventoryQueryRequest) SetInventoryType(_inventoryType int64) error {
    r._inventoryType = _inventoryType
    r.Set("inventory_type", _inventoryType)
    return nil
}

// InventoryType Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetInventoryType() int64 {
    return r._inventoryType
}
// Type Setter
// 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
func (r *TaobaoWlbWmsInventoryQueryRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetType() int64 {
    return r._type
}
// BatchCode Setter
// 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
func (r *TaobaoWlbWmsInventoryQueryRequest) SetBatchCode(_batchCode string) error {
    r._batchCode = _batchCode
    r.Set("batch_code", _batchCode)
    return nil
}

// BatchCode Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetBatchCode() string {
    return r._batchCode
}
// ProduceDate Setter
// 生产日期，type=2时字段传值有效。
func (r *TaobaoWlbWmsInventoryQueryRequest) SetProduceDate(_produceDate string) error {
    r._produceDate = _produceDate
    r.Set("produce_date", _produceDate)
    return nil
}

// ProduceDate Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetProduceDate() string {
    return r._produceDate
}
// DueDate Setter
// 失效日期，type=2时字段传值有效。
func (r *TaobaoWlbWmsInventoryQueryRequest) SetDueDate(_dueDate string) error {
    r._dueDate = _dueDate
    r.Set("due_date", _dueDate)
    return nil
}

// DueDate Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetDueDate() string {
    return r._dueDate
}
// ChannelCode Setter
// 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
func (r *TaobaoWlbWmsInventoryQueryRequest) SetChannelCode(_channelCode string) error {
    r._channelCode = _channelCode
    r.Set("channel_code", _channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetChannelCode() string {
    return r._channelCode
}
// PageNo Setter
// 分页的第几页
func (r *TaobaoWlbWmsInventoryQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页多少条，最大50条
func (r *TaobaoWlbWmsInventoryQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
