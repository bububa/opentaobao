package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtBranddetailQueryAPIRequest 品牌详情查询 API请求
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
type TmallNrtBranddetailQueryAPIRequest struct {
	model.Params
	// 品牌id
	_brandId int64
}

// NewTmallNrtBranddetailQueryRequest 初始化TmallNrtBranddetailQueryAPIRequest对象
func NewTmallNrtBranddetailQueryRequest() *TmallNrtBranddetailQueryAPIRequest {
	return &TmallNrtBranddetailQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtBranddetailQueryAPIRequest) Reset() {
	r._brandId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtBranddetailQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.branddetail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtBranddetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtBranddetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandId is BrandId Setter
// 品牌id
func (r *TmallNrtBranddetailQueryAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallNrtBranddetailQueryAPIRequest) GetBrandId() int64 {
	return r._brandId
}

var poolTmallNrtBranddetailQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtBranddetailQueryRequest()
	},
}

// GetTmallNrtBranddetailQueryRequest 从 sync.Pool 获取 TmallNrtBranddetailQueryAPIRequest
func GetTmallNrtBranddetailQueryAPIRequest() *TmallNrtBranddetailQueryAPIRequest {
	return poolTmallNrtBranddetailQueryAPIRequest.Get().(*TmallNrtBranddetailQueryAPIRequest)
}

// ReleaseTmallNrtBranddetailQueryAPIRequest 将 TmallNrtBranddetailQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrtBranddetailQueryAPIRequest(v *TmallNrtBranddetailQueryAPIRequest) {
	v.Reset()
	poolTmallNrtBranddetailQueryAPIRequest.Put(v)
}
