package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
非交易库存调整单 APIRequest
taobao.inventory.adjust.external

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家非交易调整库存，调拨出库、盘点等时调用
*/
type TaobaoInventoryAdjustExternalRequest struct {
    model.Params

    // test
    reduceType   string 

    // 外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
    bizType   string 

    // test
    operateType   string 

    // 商家外部定单号
    bizUniqueCode   string 

    // 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]
    items   string 

    // 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
    occupyOperateCode   string 

    // 业务操作时间
    operateTime   string 

    // 商家仓库编码
    storeCode   string 

}

func NewTaobaoInventoryAdjustExternalRequest() *TaobaoInventoryAdjustExternalRequest{
    return &TaobaoInventoryAdjustExternalRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryAdjustExternalRequest) GetApiMethodName() string {
    return "taobao.inventory.adjust.external"
}

func (r TaobaoInventoryAdjustExternalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryAdjustExternalRequest) SetReduceType(reduceType string) error {
    r.reduceType = reduceType
    r.Set("reduce_type", reduceType)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetReduceType() string {
    return r.reduceType
}

func (r *TaobaoInventoryAdjustExternalRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetBizType() string {
    return r.bizType
}

func (r *TaobaoInventoryAdjustExternalRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetOperateType() string {
    return r.operateType
}

func (r *TaobaoInventoryAdjustExternalRequest) SetBizUniqueCode(bizUniqueCode string) error {
    r.bizUniqueCode = bizUniqueCode
    r.Set("biz_unique_code", bizUniqueCode)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetBizUniqueCode() string {
    return r.bizUniqueCode
}

func (r *TaobaoInventoryAdjustExternalRequest) SetItems(items string) error {
    r.items = items
    r.Set("items", items)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetItems() string {
    return r.items
}

func (r *TaobaoInventoryAdjustExternalRequest) SetOccupyOperateCode(occupyOperateCode string) error {
    r.occupyOperateCode = occupyOperateCode
    r.Set("occupy_operate_code", occupyOperateCode)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetOccupyOperateCode() string {
    return r.occupyOperateCode
}

func (r *TaobaoInventoryAdjustExternalRequest) SetOperateTime(operateTime string) error {
    r.operateTime = operateTime
    r.Set("operate_time", operateTime)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetOperateTime() string {
    return r.operateTime
}

func (r *TaobaoInventoryAdjustExternalRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoInventoryAdjustExternalRequest) GetStoreCode() string {
    return r.storeCode
}

