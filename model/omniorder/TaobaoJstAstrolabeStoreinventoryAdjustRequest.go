package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
后端商品库存占用调整接口 API请求
taobao.jst.astrolabe.storeinventory.adjust

当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
*/
type TaobaoJstAstrolabeStoreinventoryAdjustRequest struct {
    model.Params
    // 操作时间
    operationTime   string
    // 库存调整信息
    inventoryAdjustInfo   *InventoryAdjustInfo
}

// 初始化TaobaoJstAstrolabeStoreinventoryAdjustRequest对象
func NewTaobaoJstAstrolabeStoreinventoryAdjustRequest() *TaobaoJstAstrolabeStoreinventoryAdjustRequest{
    return &TaobaoJstAstrolabeStoreinventoryAdjustRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryAdjustRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.adjust"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryAdjustRequest) SetOperationTime(operationTime string) error {
    r.operationTime = operationTime
    r.Set("operation_time", operationTime)
    return nil
}

// OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryAdjustRequest) GetOperationTime() string {
    return r.operationTime
}
// InventoryAdjustInfo Setter
// 库存调整信息
func (r *TaobaoJstAstrolabeStoreinventoryAdjustRequest) SetInventoryAdjustInfo(inventoryAdjustInfo *InventoryAdjustInfo) error {
    r.inventoryAdjustInfo = inventoryAdjustInfo
    r.Set("inventory_adjust_info", inventoryAdjustInfo)
    return nil
}

// InventoryAdjustInfo Getter
func (r TaobaoJstAstrolabeStoreinventoryAdjustRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
    return r.inventoryAdjustInfo
}
