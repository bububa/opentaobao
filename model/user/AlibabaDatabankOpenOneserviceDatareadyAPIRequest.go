package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadatabankopenoneservicedatareadyAPIRequest 瓴羊DaaS消费者增长CGP查询DataReady API请求
// alibaba.databank.open.oneservice.dataready
//
// 瓴羊DaaS消费者增长CGP取数接口
type AlibabadatabankopenoneservicedatareadyAPIRequest struct {
	model.Params
	// dataReady类型
	_dataReadyTypes string
	// 品牌ID
	_brandId int64
}

// NewAlibabadatabankopenoneservicedatareadyRequest 初始化AlibabadatabankopenoneservicedatareadyAPIRequest对象
func NewAlibabadatabankopenoneservicedatareadyRequest() *AlibabadatabankopenoneservicedatareadyAPIRequest {
	return &AlibabadatabankopenoneservicedatareadyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadatabankopenoneservicedatareadyAPIRequest) GetApiMethodName() string {
	return "alibaba.databank.open.oneservice.dataready"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadatabankopenoneservicedatareadyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadatabankopenoneservicedatareadyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataReadyTypes is DataReadyTypes Setter
// dataReady类型
func (r *AlibabadatabankopenoneservicedatareadyAPIRequest) SetDataReadyTypes(_dataReadyTypes string) error {
	r._dataReadyTypes = _dataReadyTypes
	r.Set("data_ready_types", _dataReadyTypes)
	return nil
}

// GetDataReadyTypes DataReadyTypes Getter
func (r AlibabadatabankopenoneservicedatareadyAPIRequest) GetDataReadyTypes() string {
	return r._dataReadyTypes
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *AlibabadatabankopenoneservicedatareadyAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r AlibabadatabankopenoneservicedatareadyAPIRequest) GetBrandId() int64 {
	return r._brandId
}
