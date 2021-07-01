package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
派样商品库存变更同步接口 API请求
alibaba.wdk.bm.paiyang.supply.publish

淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
*/
type AlibabaWdkBmPaiyangSupplyPublishAPIRequest struct {
    model.Params
    // 请求入参
    _isvSupplySyncParam   *IsvSupplySyncParam
}

// 初始化AlibabaWdkBmPaiyangSupplyPublishAPIRequest对象
func NewAlibabaWdkBmPaiyangSupplyPublishRequest() *AlibabaWdkBmPaiyangSupplyPublishAPIRequest{
    return &AlibabaWdkBmPaiyangSupplyPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.paiyang.supply.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvSupplySyncParam Setter
// 请求入参
func (r *AlibabaWdkBmPaiyangSupplyPublishAPIRequest) SetIsvSupplySyncParam(_isvSupplySyncParam *IsvSupplySyncParam) error {
    r._isvSupplySyncParam = _isvSupplySyncParam
    r.Set("isv_supply_sync_param", _isvSupplySyncParam)
    return nil
}

// IsvSupplySyncParam Getter
func (r AlibabaWdkBmPaiyangSupplyPublishAPIRequest) GetIsvSupplySyncParam() *IsvSupplySyncParam {
    return r._isvSupplySyncParam
}
