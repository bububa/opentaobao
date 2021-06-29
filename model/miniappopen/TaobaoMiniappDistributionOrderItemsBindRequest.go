package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-基于投放计划绑定/解绑商品 API请求
taobao.miniapp.distribution.order.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
*/
type TaobaoMiniappDistributionOrderItemsBindRequest struct {
    model.Params
    // 商品id列表
    _targetEntityList   []string
    // true表示新增绑定，false表示解绑
    _addBind   bool
    // 投放计划标识id
    _distributeId   int64
}

// 初始化TaobaoMiniappDistributionOrderItemsBindRequest对象
func NewTaobaoMiniappDistributionOrderItemsBindRequest() *TaobaoMiniappDistributionOrderItemsBindRequest{
    return &TaobaoMiniappDistributionOrderItemsBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.items.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TargetEntityList Setter
// 商品id列表
func (r *TaobaoMiniappDistributionOrderItemsBindRequest) SetTargetEntityList(_targetEntityList []string) error {
    r._targetEntityList = _targetEntityList
    r.Set("target_entity_list", _targetEntityList)
    return nil
}

// TargetEntityList Getter
func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetTargetEntityList() []string {
    return r._targetEntityList
}
// AddBind Setter
// true表示新增绑定，false表示解绑
func (r *TaobaoMiniappDistributionOrderItemsBindRequest) SetAddBind(_addBind bool) error {
    r._addBind = _addBind
    r.Set("add_bind", _addBind)
    return nil
}

// AddBind Getter
func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetAddBind() bool {
    return r._addBind
}
// DistributeId Setter
// 投放计划标识id
func (r *TaobaoMiniappDistributionOrderItemsBindRequest) SetDistributeId(_distributeId int64) error {
    r._distributeId = _distributeId
    r.Set("distribute_id", _distributeId)
    return nil
}

// DistributeId Getter
func (r TaobaoMiniappDistributionOrderItemsBindRequest) GetDistributeId() int64 {
    return r._distributeId
}
