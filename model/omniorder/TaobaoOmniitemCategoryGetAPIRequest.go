package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniitemCategoryGetAPIRequest 全渠道商品轻发布类目信息 API请求
// taobao.omniitem.category.get
//
// 全渠道商品轻发布类目信息
type TaobaoOmniitemCategoryGetAPIRequest struct {
	model.Params
	// 全渠道商品类目ID，不填表示获取所有全渠道商品类目信息
	_categoryId int64
}

// NewTaobaoOmniitemCategoryGetRequest 初始化TaobaoOmniitemCategoryGetAPIRequest对象
func NewTaobaoOmniitemCategoryGetRequest() *TaobaoOmniitemCategoryGetAPIRequest {
	return &TaobaoOmniitemCategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemCategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemCategoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCategoryId is CategoryId Setter
// 全渠道商品类目ID，不填表示获取所有全渠道商品类目信息
func (r *TaobaoOmniitemCategoryGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaoOmniitemCategoryGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
