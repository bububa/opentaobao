package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemqualificationdisplaygetAPIRequest 资质采集配置异步获取接口 API请求
// taobao.item.qualification.display.get
//
// 根据类目，商品，属性等参与动态获得资质采集配置
type TaobaoitemqualificationdisplaygetAPIRequest struct {
	model.Params
	// 参数列表，为key和value都是string的map的转化的json格式
	_param string
	// 商品id
	_itemId int64
	// 类目id
	_categoryId int64
}

// NewTaobaoitemqualificationdisplaygetRequest 初始化TaobaoitemqualificationdisplaygetAPIRequest对象
func NewTaobaoitemqualificationdisplaygetRequest() *TaobaoitemqualificationdisplaygetAPIRequest {
	return &TaobaoitemqualificationdisplaygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemqualificationdisplaygetAPIRequest) GetApiMethodName() string {
	return "taobao.item.qualification.display.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemqualificationdisplaygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemqualificationdisplaygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数列表，为key和value都是string的map的转化的json格式
func (r *TaobaoitemqualificationdisplaygetAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoitemqualificationdisplaygetAPIRequest) GetParam() string {
	return r._param
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoitemqualificationdisplaygetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoitemqualificationdisplaygetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCategoryId is CategoryId Setter
// 类目id
func (r *TaobaoitemqualificationdisplaygetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaoitemqualificationdisplaygetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
