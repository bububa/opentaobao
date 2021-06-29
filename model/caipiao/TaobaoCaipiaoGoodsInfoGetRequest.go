package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id与appkey获取商品信息 API请求
taobao.caipiao.goods.info.get

根据卖家id与appkey获取商品信息。
*/
type TaobaoCaipiaoGoodsInfoGetRequest struct {
    model.Params
}

// 初始化TaobaoCaipiaoGoodsInfoGetRequest对象
func NewTaobaoCaipiaoGoodsInfoGetRequest() *TaobaoCaipiaoGoodsInfoGetRequest{
    return &TaobaoCaipiaoGoodsInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoGoodsInfoGetRequest) GetApiMethodName() string {
    return "taobao.caipiao.goods.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoGoodsInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
