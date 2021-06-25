package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟商品库存查询 APIRequest
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

func NewTaobaoWlbWmsInventoryQueryRequest() *TaobaoWlbWmsInventoryQueryRequest{
    return &TaobaoWlbWmsInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.query"
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsInventoryQueryRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetInventoryType(inventoryType int64) error {
    r.inventoryType = inventoryType
    r.Set("inventory_type", inventoryType)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetInventoryType() int64 {
    return r.inventoryType
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetBatchCode(batchCode string) error {
    r.batchCode = batchCode
    r.Set("batch_code", batchCode)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetBatchCode() string {
    return r.batchCode
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetProduceDate(produceDate string) error {
    r.produceDate = produceDate
    r.Set("produce_date", produceDate)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetProduceDate() string {
    return r.produceDate
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetDueDate(dueDate string) error {
    r.dueDate = dueDate
    r.Set("due_date", dueDate)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetDueDate() string {
    return r.dueDate
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetChannelCode() string {
    return r.channelCode
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbWmsInventoryQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbWmsInventoryQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

