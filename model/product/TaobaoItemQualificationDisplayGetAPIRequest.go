package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemQualificationDisplayGetAPIRequest 资质采集配置异步获取接口 API请求
// taobao.item.qualification.display.get
//
// 根据类目，商品，属性等参与动态获得资质采集配置
type TaobaoItemQualificationDisplayGetAPIRequest struct {
	model.Params
	// 参数列表，为key和value都是string的map的转化的json格式
	_param string
	// 商品id
	_itemId int64
	// 类目id
	_categoryId int64
}

// NewTaobaoItemQualificationDisplayGetRequest 初始化TaobaoItemQualificationDisplayGetAPIRequest对象
func NewTaobaoItemQualificationDisplayGetRequest() *TaobaoItemQualificationDisplayGetAPIRequest {
	return &TaobaoItemQualificationDisplayGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.qualification.display.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 参数列表，为key和value都是string的map的转化的json格式
func (r *TaobaoItemQualificationDisplayGetAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetParam() string {
	return r._param
}

// Set is ItemId Setter
// 商品id
func (r *TaobaoItemQualificationDisplayGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is CategoryId Setter
// 类目id
func (r *TaobaoItemQualificationDisplayGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// Get CategoryId Getter
func (r TaobaoItemQualificationDisplayGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
