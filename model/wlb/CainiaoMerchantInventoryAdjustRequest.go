package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家库存调整 API请求
cainiao.merchant.inventory.adjust

商家仓库存调整接口，目前仅支持全量更新
*/
type CainiaoMerchantInventoryAdjustAPIRequest struct {
    model.Params
    // 商家仓编辑库存
    _adjustRequest   []MerStoreInvAdjustDTO
    // 调用方应用名
    _appName   string
    // 操作
    _operation   string
}

// 初始化CainiaoMerchantInventoryAdjustAPIRequest对象
func NewCainiaoMerchantInventoryAdjustRequest() *CainiaoMerchantInventoryAdjustAPIRequest{
    return &CainiaoMerchantInventoryAdjustAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetApiMethodName() string {
    return "cainiao.merchant.inventory.adjust"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdjustRequest Setter
// 商家仓编辑库存
func (r *CainiaoMerchantInventoryAdjustAPIRequest) SetAdjustRequest(_adjustRequest []MerStoreInvAdjustDTO) error {
    r._adjustRequest = _adjustRequest
    r.Set("adjust_request", _adjustRequest)
    return nil
}

// AdjustRequest Getter
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetAdjustRequest() []MerStoreInvAdjustDTO {
    return r._adjustRequest
}
// AppName Setter
// 调用方应用名
func (r *CainiaoMerchantInventoryAdjustAPIRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetAppName() string {
    return r._appName
}
// Operation Setter
// 操作
func (r *CainiaoMerchantInventoryAdjustAPIRequest) SetOperation(_operation string) error {
    r._operation = _operation
    r.Set("operation", _operation)
    return nil
}

// Operation Getter
func (r CainiaoMerchantInventoryAdjustAPIRequest) GetOperation() string {
    return r._operation
}
