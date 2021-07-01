package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家被授权品牌列表和类目列表 API请求
taobao.itemcats.authorize.get

查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
*/
type TaobaoItemcatsAuthorizeGetAPIRequest struct {
    model.Params
    // 需要返回的字段。目前支持有：<br/>brand.vid, brand.name, <br/>item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,<br/>xinpin_item_cat.cid, <br/>xinpin_item_cat.name, <br/>xinpin_item_cat.status,<br/>xinpin_item_cat.sort_order,<br/>xinpin_item_cat.parent_cid,<br/>xinpin_item_cat.is_parent
    _fields   []string
}

// 初始化TaobaoItemcatsAuthorizeGetAPIRequest对象
func NewTaobaoItemcatsAuthorizeGetRequest() *TaobaoItemcatsAuthorizeGetAPIRequest{
    return &TaobaoItemcatsAuthorizeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetApiMethodName() string {
    return "taobao.itemcats.authorize.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段。目前支持有：<br/>brand.vid, brand.name, <br/>item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,<br/>xinpin_item_cat.cid, <br/>xinpin_item_cat.name, <br/>xinpin_item_cat.status,<br/>xinpin_item_cat.sort_order,<br/>xinpin_item_cat.parent_cid,<br/>xinpin_item_cat.is_parent
func (r *TaobaoItemcatsAuthorizeGetAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetFields() []string {
    return r._fields
}
