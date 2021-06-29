package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存占用调整接口 API请求
taobao.jst.astrolabe.storeinventory.itemadjust

当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
*/
type TaobaoJstAstrolabeStoreinventoryItemadjustRequest struct {
    model.Params
    // 操作时间
    _operationTime   string
    // 库存调整信息
    _inventoryAdjustInfo   *InventoryAdjustInfo
}

// 初始化TaobaoJstAstrolabeStoreinventoryItemadjustRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemadjustRequest() *TaobaoJstAstrolabeStoreinventoryItemadjustRequest{
    return &TaobaoJstAstrolabeStoreinventoryItemadjustRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.itemadjust"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryItemadjustRequest) SetOperationTime(_operationTime string) error {
    r._operationTime = _operationTime
    r.Set("operation_time", _operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetOperationTime() string {
    return r._operationTime
}
// InventoryAdjustInfo Setter
// 库存调整信息
func (r *TaobaoJstAstrolabeStoreinventoryItemadjustRequest) SetInventoryAdjustInfo(_inventoryAdjustInfo *InventoryAdjustInfo) error {
    r._inventoryAdjustInfo = _inventoryAdjustInfo
    r.Set("inventory_adjust_info", _inventoryAdjustInfo)
    return nil
}

// InventoryAdjustInfo Getter
func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
    return r._inventoryAdjustInfo
}
