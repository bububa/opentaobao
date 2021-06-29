package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-基于投放计划绑定/解绑全店商品 API请求
taobao.miniapp.distribution.order.items.all.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。该接口对应的能力是全店投放，该能力的使用，需要联系平台运营进行人工申请，申请通过以后方可调用成功。
*/
type TaobaoMiniappDistributionOrderItemsAllBindRequest struct {
    model.Params
    // 绑定/解绑的入参信息
    _allItemBindRequest   *DistributionOrderBindTargetEntityOpenRequestV2
}

// 初始化TaobaoMiniappDistributionOrderItemsAllBindRequest对象
func NewTaobaoMiniappDistributionOrderItemsAllBindRequest() *TaobaoMiniappDistributionOrderItemsAllBindRequest{
    return &TaobaoMiniappDistributionOrderItemsAllBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderItemsAllBindRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.items.all.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderItemsAllBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AllItemBindRequest Setter
// 绑定/解绑的入参信息
func (r *TaobaoMiniappDistributionOrderItemsAllBindRequest) SetAllItemBindRequest(_allItemBindRequest *DistributionOrderBindTargetEntityOpenRequestV2) error {
    r._allItemBindRequest = _allItemBindRequest
    r.Set("all_item_bind_request", _allItemBindRequest)
    return nil
}

// AllItemBindRequest Getter
func (r TaobaoMiniappDistributionOrderItemsAllBindRequest) GetAllItemBindRequest() *DistributionOrderBindTargetEntityOpenRequestV2 {
    return r._allItemBindRequest
}
