package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercatslistgetAPIRequest 获取前台展示的店铺内卖家自定义商品类目 API请求
// taobao.sellercats.list.get
//
// 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
type TaobaosellercatslistgetAPIRequest struct {
	model.Params
}

// NewTaobaosellercatslistgetRequest 初始化TaobaosellercatslistgetAPIRequest对象
func NewTaobaosellercatslistgetRequest() *TaobaosellercatslistgetAPIRequest {
	return &TaobaosellercatslistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercatslistgetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercatslistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercatslistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
