package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部库存变化通知（企业物流用户使用） APIRequest
taobao.wlb.out.inventory.change.notify

拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
*/
type TaobaoWlbOutInventoryChangeNotifyRequest struct {
    model.Params

    // WLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku
    type   string 

    // OUT--出库 IN--入库
    opType   string 

    // （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购
    source   string 

    // 物流宝商品id或前台宝贝id（由type类型决定）
    itemId   int64 

    // 库存变化数量
    changeCount   int64 

    // 本次库存变化后库存余额
    resultCount   int64 

    // 订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号
    orderSourceCode   string 

    // 库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。
    outBizCode   string 

    // 目前非必须，系统会选择默认值
    storeCode   string 

}

func NewTaobaoWlbOutInventoryChangeNotifyRequest() *TaobaoWlbOutInventoryChangeNotifyRequest{
    return &TaobaoWlbOutInventoryChangeNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetApiMethodName() string {
    return "taobao.wlb.out.inventory.change.notify"
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetType() string {
    return r.type
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetOpType(opType string) error {
    r.opType = opType
    r.Set("op_type", opType)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetOpType() string {
    return r.opType
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetSource() string {
    return r.source
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetChangeCount(changeCount int64) error {
    r.changeCount = changeCount
    r.Set("change_count", changeCount)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetChangeCount() int64 {
    return r.changeCount
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetResultCount(resultCount int64) error {
    r.resultCount = resultCount
    r.Set("result_count", resultCount)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetResultCount() int64 {
    return r.resultCount
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetOrderSourceCode(orderSourceCode string) error {
    r.orderSourceCode = orderSourceCode
    r.Set("order_source_code", orderSourceCode)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetOrderSourceCode() string {
    return r.orderSourceCode
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetOutBizCode(outBizCode string) error {
    r.outBizCode = outBizCode
    r.Set("out_biz_code", outBizCode)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetOutBizCode() string {
    return r.outBizCode
}

func (r *TaobaoWlbOutInventoryChangeNotifyRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbOutInventoryChangeNotifyRequest) GetStoreCode() string {
    return r.storeCode
}

