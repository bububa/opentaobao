package product

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaGpuAddSchemaGetAPIRequest) Reset() {
	r._leafCatId = 0
	r._brandId = 0
	r._providerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuAddSchemaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuAddSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGpuAddSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLeafCatId is LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuAddSchemaGetAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// GetLeafCatId LeafCatId Getter
func (r AlibabaGpuAddSchemaGetAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *AlibabaGpuAddSchemaGetAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r AlibabaGpuAddSchemaGetAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuAddSchemaGetAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaGpuAddSchemaGetAPIRequest) GetProviderId() int64 {
	return r._providerId
}

var poolAlibabaGpuAddSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaGpuAddSchemaGetRequest()
	},
}

// GetAlibabaGpuAddSchemaGetRequest 从 sync.Pool 获取 AlibabaGpuAddSchemaGetAPIRequest
func GetAlibabaGpuAddSchemaGetAPIRequest() *AlibabaGpuAddSchemaGetAPIRequest {
	return poolAlibabaGpuAddSchemaGetAPIRequest.Get().(*AlibabaGpuAddSchemaGetAPIRequest)
}

// ReleaseAlibabaGpuAddSchemaGetAPIRequest 将 AlibabaGpuAddSchemaGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaGpuAddSchemaGetAPIRequest(v *AlibabaGpuAddSchemaGetAPIRequest) {
	v.Reset()
	poolAlibabaGpuAddSchemaGetAPIRequest.Put(v)
}
