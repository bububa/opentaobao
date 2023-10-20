package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryGetAPIRequest 商品发布类目获取 API请求
// alibaba.icbu.category.get
//
// 获取商品发布类目
type AlibabaIcbuCategoryGetAPIRequest struct {
	model.Params
	// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
	_catId int64
}

// NewAlibabaIcbuCategoryGetRequest 初始化AlibabaIcbuCategoryGetAPIRequest对象
func NewAlibabaIcbuCategoryGetRequest() *AlibabaIcbuCategoryGetAPIRequest {
	return &AlibabaIcbuCategoryGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuCategoryGetAPIRequest) Reset() {
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaIcbuCategoryGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaIcbuCategoryGetAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolAlibabaIcbuCategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuCategoryGetRequest()
	},
}

// GetAlibabaIcbuCategoryGetRequest 从 sync.Pool 获取 AlibabaIcbuCategoryGetAPIRequest
func GetAlibabaIcbuCategoryGetAPIRequest() *AlibabaIcbuCategoryGetAPIRequest {
	return poolAlibabaIcbuCategoryGetAPIRequest.Get().(*AlibabaIcbuCategoryGetAPIRequest)
}

// ReleaseAlibabaIcbuCategoryGetAPIRequest 将 AlibabaIcbuCategoryGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuCategoryGetAPIRequest(v *AlibabaIcbuCategoryGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuCategoryGetAPIRequest.Put(v)
}
