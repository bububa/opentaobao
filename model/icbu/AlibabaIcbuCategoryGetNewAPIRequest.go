package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryGetNewAPIRequest （新）ICBU类目树获取接口 API请求
// alibaba.icbu.category.get.new
//
// 获取商品发布类目
type AlibabaIcbuCategoryGetNewAPIRequest struct {
	model.Params
	// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
	_catId int64
}

// NewAlibabaIcbuCategoryGetNewRequest 初始化AlibabaIcbuCategoryGetNewAPIRequest对象
func NewAlibabaIcbuCategoryGetNewRequest() *AlibabaIcbuCategoryGetNewAPIRequest {
	return &AlibabaIcbuCategoryGetNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuCategoryGetNewAPIRequest) Reset() {
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.get.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaIcbuCategoryGetNewAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaIcbuCategoryGetNewAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolAlibabaIcbuCategoryGetNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuCategoryGetNewRequest()
	},
}

// GetAlibabaIcbuCategoryGetNewRequest 从 sync.Pool 获取 AlibabaIcbuCategoryGetNewAPIRequest
func GetAlibabaIcbuCategoryGetNewAPIRequest() *AlibabaIcbuCategoryGetNewAPIRequest {
	return poolAlibabaIcbuCategoryGetNewAPIRequest.Get().(*AlibabaIcbuCategoryGetNewAPIRequest)
}

// ReleaseAlibabaIcbuCategoryGetNewAPIRequest 将 AlibabaIcbuCategoryGetNewAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuCategoryGetNewAPIRequest(v *AlibabaIcbuCategoryGetNewAPIRequest) {
	v.Reset()
	poolAlibabaIcbuCategoryGetNewAPIRequest.Put(v)
}
