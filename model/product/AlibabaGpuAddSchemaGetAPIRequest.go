package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuAddSchemaGetAPIRequest 获取产品发布规则接口 API请求
// alibaba.gpu.add.schema.get
//
// 获取产品发布规则接口
type AlibabaGpuAddSchemaGetAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 品牌ID
	_brandId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabaGpuAddSchemaGetRequest 初始化AlibabaGpuAddSchemaGetAPIRequest对象
func NewAlibabaGpuAddSchemaGetRequest() *AlibabaGpuAddSchemaGetAPIRequest {
	return &AlibabaGpuAddSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuAddSchemaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuAddSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuAddSchemaGetAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// Get LeafCatId Getter
func (r AlibabaGpuAddSchemaGetAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// Set is BrandId Setter
// 品牌ID
func (r *AlibabaGpuAddSchemaGetAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// Get BrandId Getter
func (r AlibabaGpuAddSchemaGetAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// Set is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuAddSchemaGetAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// Get ProviderId Getter
func (r AlibabaGpuAddSchemaGetAPIRequest) GetProviderId() int64 {
	return r._providerId
}
