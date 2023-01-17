package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintIsvtemplatesGetAPIRequest 获取商家使用的标准模板 API请求
// cainiao.cloudprint.isvtemplates.get
//
// 获取商家使用的标准模板
type CainiaoCloudprintIsvtemplatesGetAPIRequest struct {
	model.Params
}

// NewCainiaoCloudprintIsvtemplatesGetRequest 初始化CainiaoCloudprintIsvtemplatesGetAPIRequest对象
func NewCainiaoCloudprintIsvtemplatesGetRequest() *CainiaoCloudprintIsvtemplatesGetAPIRequest {
	return &CainiaoCloudprintIsvtemplatesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintIsvtemplatesGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.isvtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintIsvtemplatesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintIsvtemplatesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
