package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-商品绑定/解绑 APIRequest
taobao.miniapp.distribution.items.bind

提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。
*/
type TaobaoMiniappDistributionItemsBindRequest struct {
    model.Params

    // 商品id列表
    targetEntityList   []string 

    // 投放的商家应用完整链接
    url   string 

    // true表示新增绑定，false表示解绑
    addBind   bool 

}

func NewTaobaoMiniappDistributionItemsBindRequest() *TaobaoMiniappDistributionItemsBindRequest{
    return &TaobaoMiniappDistributionItemsBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappDistributionItemsBindRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.items.bind"
}

func (r TaobaoMiniappDistributionItemsBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappDistributionItemsBindRequest) SetTargetEntityList(targetEntityList []string) error {
    r.targetEntityList = targetEntityList
    r.Set("target_entity_list", targetEntityList)
    return nil
}

func (r TaobaoMiniappDistributionItemsBindRequest) GetTargetEntityList() []string {
    return r.targetEntityList
}

func (r *TaobaoMiniappDistributionItemsBindRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoMiniappDistributionItemsBindRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoMiniappDistributionItemsBindRequest) SetAddBind(addBind bool) error {
    r.addBind = addBind
    r.Set("add_bind", addBind)
    return nil
}

func (r TaobaoMiniappDistributionItemsBindRequest) GetAddBind() bool {
    return r.addBind
}

