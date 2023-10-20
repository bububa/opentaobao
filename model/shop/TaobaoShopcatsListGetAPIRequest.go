package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoshopcatslistgetAPIRequest 获取前台展示的店铺类目 API请求
// taobao.shopcats.list.get
//
// 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoshopcatslistgetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
	_fields []string
}

// NewTaobaoshopcatslistgetRequest 初始化TaobaoshopcatslistgetAPIRequest对象
func NewTaobaoshopcatslistgetRequest() *TaobaoshopcatslistgetAPIRequest {
	return &TaobaoshopcatslistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoshopcatslistgetAPIRequest) GetApiMethodName() string {
	return "taobao.shopcats.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoshopcatslistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoshopcatslistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
func (r *TaobaoshopcatslistgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoshopcatslistgetAPIRequest) GetFields() []string {
	return r._fields
}
