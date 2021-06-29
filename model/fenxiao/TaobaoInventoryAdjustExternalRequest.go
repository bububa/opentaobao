package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
非交易库存调整单 API请求
taobao.inventory.adjust.external

建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
商家非交易调整库存，调拨出库、盘点等时调用
*/
type TaobaoInventoryAdjustExternalRequest struct {
    model.Params
    // test
    _reduceType   string
    // 外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
    _bizType   string
    // test
    _operateType   string
    // 商家外部定单号
    _bizUniqueCode   string
    // 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]
    _items   string
    // 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
    _occupyOperateCode   string
    // 业务操作时间
    _operateTime   string
    // 商家仓库编码
    _storeCode   string
}

// 初始化TaobaoInventoryAdjustExternalRequest对象
func NewTaobaoInventoryAdjustExternalRequest() *TaobaoInventoryAdjustExternalRequest{
    return &TaobaoInventoryAdjustExternalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryAdjustExternalRequest) GetApiMethodName() string {
    return "taobao.inventory.adjust.external"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryAdjustExternalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReduceType Setter
// test
func (r *TaobaoInventoryAdjustExternalRequest) SetReduceType(_reduceType string) error {
    r._reduceType = _reduceType
    r.Set("reduce_type", _reduceType)
    return nil
}

// ReduceType Getter
func (r TaobaoInventoryAdjustExternalRequest) GetReduceType() string {
    return r._reduceType
}
// BizType Setter
// 外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
func (r *TaobaoInventoryAdjustExternalRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TaobaoInventoryAdjustExternalRequest) GetBizType() string {
    return r._bizType
}
// OperateType Setter
// test
func (r *TaobaoInventoryAdjustExternalRequest) SetOperateType(_operateType string) error {
    r._operateType = _operateType
    r.Set("operate_type", _operateType)
    return nil
}

// OperateType Getter
func (r TaobaoInventoryAdjustExternalRequest) GetOperateType() string {
    return r._operateType
}
// BizUniqueCode Setter
// 商家外部定单号
func (r *TaobaoInventoryAdjustExternalRequest) SetBizUniqueCode(_bizUniqueCode string) error {
    r._bizUniqueCode = _bizUniqueCode
    r.Set("biz_unique_code", _bizUniqueCode)
    return nil
}

// BizUniqueCode Getter
func (r TaobaoInventoryAdjustExternalRequest) GetBizUniqueCode() string {
    return r._bizUniqueCode
}
// Items Setter
// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]
func (r *TaobaoInventoryAdjustExternalRequest) SetItems(_items string) error {
    r._items = _items
    r.Set("items", _items)
    return nil
}

// Items Getter
func (r TaobaoInventoryAdjustExternalRequest) GetItems() string {
    return r._items
}
// OccupyOperateCode Setter
// 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
func (r *TaobaoInventoryAdjustExternalRequest) SetOccupyOperateCode(_occupyOperateCode string) error {
    r._occupyOperateCode = _occupyOperateCode
    r.Set("occupy_operate_code", _occupyOperateCode)
    return nil
}

// OccupyOperateCode Getter
func (r TaobaoInventoryAdjustExternalRequest) GetOccupyOperateCode() string {
    return r._occupyOperateCode
}
// OperateTime Setter
// 业务操作时间
func (r *TaobaoInventoryAdjustExternalRequest) SetOperateTime(_operateTime string) error {
    r._operateTime = _operateTime
    r.Set("operate_time", _operateTime)
    return nil
}

// OperateTime Getter
func (r TaobaoInventoryAdjustExternalRequest) GetOperateTime() string {
    return r._operateTime
}
// StoreCode Setter
// 商家仓库编码
func (r *TaobaoInventoryAdjustExternalRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TaobaoInventoryAdjustExternalRequest) GetStoreCode() string {
    return r._storeCode
}
