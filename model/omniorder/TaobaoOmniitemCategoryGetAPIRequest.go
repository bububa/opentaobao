package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemcategorygetAPIRequest 全渠道商品轻发布类目信息 API请求
// taobao.omniitem.category.get
//
// 全渠道商品轻发布类目信息
type TaobaoomniitemcategorygetAPIRequest struct {
	model.Params
	// 全渠道商品类目ID，不填表示获取所有全渠道商品类目信息
	_categoryId int64
}

// NewTaobaoomniitemcategorygetRequest 初始化TaobaoomniitemcategorygetAPIRequest对象
func NewTaobaoomniitemcategorygetRequest() *TaobaoomniitemcategorygetAPIRequest {
	return &TaobaoomniitemcategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniitemcategorygetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniitemcategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniitemcategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 全渠道商品类目ID，不填表示获取所有全渠道商品类目信息
func (r *TaobaoomniitemcategorygetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TaobaoomniitemcategorygetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
