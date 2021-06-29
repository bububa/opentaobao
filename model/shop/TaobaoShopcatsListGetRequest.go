package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取前台展示的店铺类目 API请求
taobao.shopcats.list.get

获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
*/
type TaobaoShopcatsListGetRequest struct {
    model.Params
    // 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
    fields   []string
}

// 初始化TaobaoShopcatsListGetRequest对象
func NewTaobaoShopcatsListGetRequest() *TaobaoShopcatsListGetRequest{
    return &TaobaoShopcatsListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoShopcatsListGetRequest) GetApiMethodName() string {
    return "taobao.shopcats.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoShopcatsListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
func (r *TaobaoShopcatsListGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoShopcatsListGetRequest) GetFields() []string {
    return r.fields
}
