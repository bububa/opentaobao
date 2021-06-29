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
type AlibabaWdkBmPaiyangSupplyPublishRequest struct {
    model.Params
    // 请求入参
    isvSupplySyncParam   *IsvSupplySyncParam
}

// 初始化AlibabaWdkBmPaiyangSupplyPublishRequest对象
func NewAlibabaWdkBmPaiyangSupplyPublishRequest() *AlibabaWdkBmPaiyangSupplyPublishRequest{
    return &AlibabaWdkBmPaiyangSupplyPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangSupplyPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.paiyang.supply.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangSupplyPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvSupplySyncParam Setter
// 请求入参
func (r *AlibabaWdkBmPaiyangSupplyPublishRequest) SetIsvSupplySyncParam(isvSupplySyncParam *IsvSupplySyncParam) error {
    r.isvSupplySyncParam = isvSupplySyncParam
    r.Set("isv_supply_sync_param", isvSupplySyncParam)
    return nil
}

// IsvSupplySyncParam Getter
func (r AlibabaWdkBmPaiyangSupplyPublishRequest) GetIsvSupplySyncParam() *IsvSupplySyncParam {
    return r.isvSupplySyncParam
}
