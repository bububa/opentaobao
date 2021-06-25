package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家库存调整 APIRequest
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新
*/
type CainiaoMerchantInventoryAdjustRequest struct {
    model.Params

    // 商家仓编辑库存
    adjustRequest   []MerStoreInvAdjustDto 

    // 调用方应用名
    appName   string 

    // 操作
    operation   string 

}

func NewCainiaoMerchantInventoryAdjustRequest() *CainiaoMerchantInventoryAdjustRequest{
    return &CainiaoMerchantInventoryAdjustRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoMerchantInventoryAdjustRequest) GetApiMethodName() string {
    return "cainiao.merchant.inventory.adjust"
}

func (r CainiaoMerchantInventoryAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoMerchantInventoryAdjustRequest) SetAdjustRequest(adjustRequest []MerStoreInvAdjustDto) error {
    r.adjustRequest = adjustRequest
    r.Set("adjust_request", adjustRequest)
    return nil
}

func (r CainiaoMerchantInventoryAdjustRequest) GetAdjustRequest() []MerStoreInvAdjustDto {
    return r.adjustRequest
}

func (r *CainiaoMerchantInventoryAdjustRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r CainiaoMerchantInventoryAdjustRequest) GetAppName() string {
    return r.appName
}

func (r *CainiaoMerchantInventoryAdjustRequest) SetOperation(operation string) error {
    r.operation = operation
    r.Set("operation", operation)
    return nil
}

func (r CainiaoMerchantInventoryAdjustRequest) GetOperation() string {
    return r.operation
}

