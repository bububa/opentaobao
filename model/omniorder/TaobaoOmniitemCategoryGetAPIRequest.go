package omniorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniitemCategoryGetAPIRequest) Reset() {
	r._categoryId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniitemCategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniitem.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniitemCategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniitemCategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoOmniitemCategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniitemCategoryGetRequest()
	},
}

// GetTaobaoOmniitemCategoryGetRequest 从 sync.Pool 获取 TaobaoOmniitemCategoryGetAPIRequest
func GetTaobaoOmniitemCategoryGetAPIRequest() *TaobaoOmniitemCategoryGetAPIRequest {
	return poolTaobaoOmniitemCategoryGetAPIRequest.Get().(*TaobaoOmniitemCategoryGetAPIRequest)
}

// ReleaseTaobaoOmniitemCategoryGetAPIRequest 将 TaobaoOmniitemCategoryGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniitemCategoryGetAPIRequest(v *TaobaoOmniitemCategoryGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniitemCategoryGetAPIRequest.Put(v)
}
