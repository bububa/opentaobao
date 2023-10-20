package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabagpuaddschemagetAPIRequest 获取产品发布规则接口 API请求
// alibaba.gpu.add.schema.get
//
// 获取产品发布规则接口
type AlibabagpuaddschemagetAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 品牌ID
	_brandId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabagpuaddschemagetRequest 初始化AlibabagpuaddschemagetAPIRequest对象
func NewAlibabagpuaddschemagetRequest() *AlibabagpuaddschemagetAPIRequest {
	return &AlibabagpuaddschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabagpuaddschemagetAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabagpuaddschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabagpuaddschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeafCatId is LeafCatId Setter
// 叶子类目ID
func (r *AlibabagpuaddschemagetAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// GetLeafCatId LeafCatId Getter
func (r AlibabagpuaddschemagetAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *AlibabagpuaddschemagetAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r AlibabagpuaddschemagetAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabagpuaddschemagetAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabagpuaddschemagetAPIRequest) GetProviderId() int64 {
	return r._providerId
}
