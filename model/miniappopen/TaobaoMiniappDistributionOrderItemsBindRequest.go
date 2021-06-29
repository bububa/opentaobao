package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-基于投放计划绑定/解绑商品 APIRequest
taobao.miniapp.distribution.order.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
*/
type TaobaoMiniappDistributionOrderItemsBindRequest struct {
    model.Params

    // 商品id列表
    targetEntityList   []string 

    // true表示新增绑定，false表示解绑
    addBind   bool 

    // 投放计划标识id
    distributeId   int64 

}

func NewTaobaoMiniappDistributionOrderItemsBindRequest() *TaobaoMiniappDistributionOrderItemsBindRequest{
    return &TaobaoMiniappDistributionOrderItemsBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.items.bind"
}

func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappDistributionOrderItemsBindRequest) SetTargetEntityList(targetEntityList []string) error {
    r.targetEntityList = targetEntityList
    r.Set("target_entity_list", targetEntityList)
    return nil
}

func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetTargetEntityList() []string {
    return r.targetEntityList
}

func (r *TaobaoMiniappDistributionOrderItemsBindRequest) SetAddBind(addBind bool) error {
    r.addBind = addBind
    r.Set("add_bind", addBind)
    return nil
}

func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetAddBind() bool {
    return r.addBind
}

func (r *TaobaoMiniappDistributionOrderItemsBindRequest) SetDistributeId(distributeId int64) error {
    r.distributeId = distributeId
    r.Set("distribute_id", distributeId)
    return nil
}

func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetDistributeId() int64 {
    return r.distributeId
}

