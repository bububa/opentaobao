package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomcnshopcatslistgetAPIRequest 店铺类目清单 API请求
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
type TaobaomcnshopcatslistgetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
	_fields []string
}

// NewTaobaomcnshopcatslistgetRequest 初始化TaobaomcnshopcatslistgetAPIRequest对象
func NewTaobaomcnshopcatslistgetRequest() *TaobaomcnshopcatslistgetAPIRequest {
	return &TaobaomcnshopcatslistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomcnshopcatslistgetAPIRequest) GetApiMethodName() string {
	return "taobao.mcn.shopcats.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomcnshopcatslistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomcnshopcatslistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
func (r *TaobaomcnshopcatslistgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaomcnshopcatslistgetAPIRequest) GetFields() []string {
	return r._fields
}
