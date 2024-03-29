package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductMatchSchemaGetAPIRequest 获取匹配产品规则 API请求
// tmall.product.match.schema.get
//
// ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallProductMatchSchemaGetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
}

// NewTmallProductMatchSchemaGetRequest 初始化TmallProductMatchSchemaGetAPIRequest对象
func NewTmallProductMatchSchemaGetRequest() *TmallProductMatchSchemaGetAPIRequest {
	return &TmallProductMatchSchemaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallProductMatchSchemaGetAPIRequest) Reset() {
	r._categoryId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductMatchSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.match.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductMatchSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductMatchSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductMatchSchemaGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallProductMatchSchemaGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

var poolTmallProductMatchSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallProductMatchSchemaGetRequest()
	},
}

// GetTmallProductMatchSchemaGetRequest 从 sync.Pool 获取 TmallProductMatchSchemaGetAPIRequest
func GetTmallProductMatchSchemaGetAPIRequest() *TmallProductMatchSchemaGetAPIRequest {
	return poolTmallProductMatchSchemaGetAPIRequest.Get().(*TmallProductMatchSchemaGetAPIRequest)
}

// ReleaseTmallProductMatchSchemaGetAPIRequest 将 TmallProductMatchSchemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallProductMatchSchemaGetAPIRequest(v *TmallProductMatchSchemaGetAPIRequest) {
	v.Reset()
	poolTmallProductMatchSchemaGetAPIRequest.Put(v)
}
