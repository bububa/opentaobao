package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-商品绑定/解绑 API请求
taobao.miniapp.distribution.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
*/
type TaobaoMiniappDistributionItemsBindRequest struct {
    model.Params
    // 商品id列表
    _targetEntityList   []string
    // 投放的商家应用完整链接
    _url   string
    // true表示新增绑定，false表示解绑
    _addBind   bool
}

// 初始化TaobaoMiniappDistributionItemsBindRequest对象
func NewTaobaoMiniappDistributionItemsBindRequest() *TaobaoMiniappDistributionItemsBindRequest{
    return &TaobaoMiniappDistributionItemsBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionItemsBindRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.items.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionItemsBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TargetEntityList Setter
// 商品id列表
func (r *TaobaoMiniappDistributionItemsBindRequest) SetTargetEntityList(_targetEntityList []string) error {
    r._targetEntityList = _targetEntityList
    r.Set("target_entity_list", _targetEntityList)
    return nil
}

// TargetEntityList Getter
func (r TaobaoMiniappDistributionItemsBindRequest) GetTargetEntityList() []string {
    return r._targetEntityList
}
// Url Setter
// 投放的商家应用完整链接
func (r *TaobaoMiniappDistributionItemsBindRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoMiniappDistributionItemsBindRequest) GetUrl() string {
    return r._url
}
// AddBind Setter
// true表示新增绑定，false表示解绑
func (r *TaobaoMiniappDistributionItemsBindRequest) SetAddBind(_addBind bool) error {
    r._addBind = _addBind
    r.Set("add_bind", _addBind)
    return nil
}

// AddBind Getter
func (r TaobaoMiniappDistributionItemsBindRequest) GetAddBind() bool {
    return r._addBind
}
