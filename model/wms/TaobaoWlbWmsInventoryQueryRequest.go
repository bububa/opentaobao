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
    itemId   string
    // 仓库编码
    storeCode   string
    // 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
    inventoryType   int64
    // 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
    type   int64
    // 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
    batchCode   string
    // 生产日期，type=2时字段传值有效。
    produceDate   string
    // 失效日期，type=2时字段传值有效。
    dueDate   string
    // 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
    channelCode   string
    // 分页的第几页
    pageNo   int64
    // 每页多少条，最大50条
    pageSize   int64
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
func (r *TaobaoWlbWmsInventoryQueryRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetItemId() string {
    return r.itemId
}
// StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsInventoryQueryRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetStoreCode() string {
    return r.storeCode
}
// InventoryType Setter
// 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
func (r *TaobaoWlbWmsInventoryQueryRequest) SetInventoryType(inventoryType int64) error {
    r.inventoryType = inventoryType
    r.Set("inventory_type", inventoryType)
    return nil
}

// InventoryType Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetInventoryType() int64 {
    return r.inventoryType
}
// Type Setter
// 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
func (r *TaobaoWlbWmsInventoryQueryRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetType() int64 {
    return r.type
}
// BatchCode Setter
// 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
func (r *TaobaoWlbWmsInventoryQueryRequest) SetBatchCode(batchCode string) error {
    r.batchCode = batchCode
    r.Set("batch_code", batchCode)
    return nil
}

// BatchCode Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetBatchCode() string {
    return r.batchCode
}
// ProduceDate Setter
// 生产日期，type=2时字段传值有效。
func (r *TaobaoWlbWmsInventoryQueryRequest) SetProduceDate(produceDate string) error {
    r.produceDate = produceDate
    r.Set("produce_date", produceDate)
    return nil
}

// ProduceDate Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetProduceDate() string {
    return r.produceDate
}
// DueDate Setter
// 失效日期，type=2时字段传值有效。
func (r *TaobaoWlbWmsInventoryQueryRequest) SetDueDate(dueDate string) error {
    r.dueDate = dueDate
    r.Set("due_date", dueDate)
    return nil
}

// DueDate Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetDueDate() string {
    return r.dueDate
}
// ChannelCode Setter
// 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
func (r *TaobaoWlbWmsInventoryQueryRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetChannelCode() string {
    return r.channelCode
}
// PageNo Setter
// 分页的第几页
func (r *TaobaoWlbWmsInventoryQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页多少条，最大50条
func (r *TaobaoWlbWmsInventoryQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWlbWmsInventoryQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
