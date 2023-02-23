package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemcatsAuthorizeGetAPIRequest 查询商家被授权品牌列表和类目列表 API请求
// taobao.itemcats.authorize.get
//
// 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
type TaobaoItemcatsAuthorizeGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：brand.vid, brand.name, item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,xinpin_item_cat.cid, xinpin_item_cat.name, xinpin_item_cat.status,xinpin_item_cat.sort_order,xinpin_item_cat.parent_cid,xinpin_item_cat.is_parent
	_fields []string
}

// NewTaobaoItemcatsAuthorizeGetRequest 初始化TaobaoItemcatsAuthorizeGetAPIRequest对象
func NewTaobaoItemcatsAuthorizeGetRequest() *TaobaoItemcatsAuthorizeGetAPIRequest {
	return &TaobaoItemcatsAuthorizeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetApiMethodName() string {
	return "taobao.itemcats.authorize.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：brand.vid, brand.name, item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,xinpin_item_cat.cid, xinpin_item_cat.name, xinpin_item_cat.status,xinpin_item_cat.sort_order,xinpin_item_cat.parent_cid,xinpin_item_cat.is_parent
func (r *TaobaoItemcatsAuthorizeGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItemcatsAuthorizeGetAPIRequest) GetFields() []string {
	return r._fields
}
