package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemrenderAPIRequest 新发商品发布页 API请求
// taobao.banamadpc.item.render
//
// 巴拿马供应商通过此接口新发商品发布页
type TaobaobanamadpcitemrenderAPIRequest struct {
	model.Params
	// 类目ID
	_catId int64
}

// NewTaobaobanamadpcitemrenderRequest 初始化TaobaobanamadpcitemrenderAPIRequest对象
func NewTaobaobanamadpcitemrenderRequest() *TaobaobanamadpcitemrenderAPIRequest {
	return &TaobaobanamadpcitemrenderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobanamadpcitemrenderAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobanamadpcitemrenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobanamadpcitemrenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 类目ID
func (r *TaobaobanamadpcitemrenderAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaobanamadpcitemrenderAPIRequest) GetCatId() int64 {
	return r._catId
}
