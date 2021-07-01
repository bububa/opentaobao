package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintStdtemplatesGetAPIRequest
获取所有的菜鸟标准电子面单模板 API请求
cainiao.cloudprint.stdtemplates.get

获取菜鸟标准电子面单模板 */
type CainiaoCloudprintStdtemplatesGetAPIRequest struct {
	model.Params
}

// NewCainiaoCloudprintStdtemplatesGetRequest 初始化CainiaoCloudprintStdtemplatesGetAPIRequest对象
func NewCainiaoCloudprintStdtemplatesGetRequest() *CainiaoCloudprintStdtemplatesGetAPIRequest {
	return &CainiaoCloudprintStdtemplatesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintStdtemplatesGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.stdtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintStdtemplatesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
