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
type TaobaoShopcatsListGetAPIRequest struct {
    model.Params
    // 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
    _fields   []string
}

// 初始化TaobaoShopcatsListGetAPIRequest对象
func NewTaobaoShopcatsListGetRequest() *TaobaoShopcatsListGetAPIRequest{
    return &TaobaoShopcatsListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoShopcatsListGetAPIRequest) GetApiMethodName() string {
    return "taobao.shopcats.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoShopcatsListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
func (r *TaobaoShopcatsListGetAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoShopcatsListGetAPIRequest) GetFields() []string {
    return r._fields
}
