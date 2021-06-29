package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存占用调整接口 APIRequest
taobao.jst.astrolabe.storeinventory.itemadjust

当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
*/
type TaobaoJstAstrolabeStoreinventoryItemadjustRequest struct {
    model.Params

    // 操作时间
    operationTime   string 

    // 库存调整信息
    inventoryAdjustInfo   *InventoryAdjustInfo 

}

func NewTaobaoJstAstrolabeStoreinventoryItemadjustRequest() *TaobaoJstAstrolabeStoreinventoryItemadjustRequest{
    return &TaobaoJstAstrolabeStoreinventoryItemadjustRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.storeinventory.itemadjust"
}

func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstAstrolabeStoreinventoryItemadjustRequest) SetOperationTime(operationTime string) error {
    r.operationTime = operationTime
    r.Set("operation_time", operationTime)
    return nil
}

func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetOperationTime() string {
    return r.operationTime
}

func (r *TaobaoJstAstrolabeStoreinventoryItemadjustRequest) SetInventoryAdjustInfo(inventoryAdjustInfo *InventoryAdjustInfo) error {
    r.inventoryAdjustInfo = inventoryAdjustInfo
    r.Set("inventory_adjust_info", inventoryAdjustInfo)
    return nil
}

func (r TaobaoJstAstrolabeStoreinventoryItemadjustRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
    return r.inventoryAdjustInfo
}

