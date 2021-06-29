package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前台展示的店铺内卖家自定义商品类目 API请求
taobao.sellercats.list.get

此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
*/
type TaobaoSellercatsListGetRequest struct {
    model.Params
}

// 初始化TaobaoSellercatsListGetRequest对象
func NewTaobaoSellercatsListGetRequest() *TaobaoSellercatsListGetRequest{
    return &TaobaoSellercatsListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercatsListGetRequest) GetApiMethodName() string {
    return "taobao.sellercats.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercatsListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
