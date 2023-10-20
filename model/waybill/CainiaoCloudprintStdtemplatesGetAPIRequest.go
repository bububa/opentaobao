package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintStdtemplatesGetAPIRequest 获取所有的菜鸟标准电子面单模板 API请求
// cainiao.cloudprint.stdtemplates.get
//
// 获取菜鸟标准电子面单模板
type CainiaoCloudprintStdtemplatesGetAPIRequest struct {
	model.Params
}

// NewCainiaoCloudprintStdtemplatesGetRequest 初始化CainiaoCloudprintStdtemplatesGetAPIRequest对象
func NewCainiaoCloudprintStdtemplatesGetRequest() *CainiaoCloudprintStdtemplatesGetAPIRequest {
	return &CainiaoCloudprintStdtemplatesGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintStdtemplatesGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintStdtemplatesGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.stdtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintStdtemplatesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintStdtemplatesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolCainiaoCloudprintStdtemplatesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintStdtemplatesGetRequest()
	},
}

// GetCainiaoCloudprintStdtemplatesGetRequest 从 sync.Pool 获取 CainiaoCloudprintStdtemplatesGetAPIRequest
func GetCainiaoCloudprintStdtemplatesGetAPIRequest() *CainiaoCloudprintStdtemplatesGetAPIRequest {
	return poolCainiaoCloudprintStdtemplatesGetAPIRequest.Get().(*CainiaoCloudprintStdtemplatesGetAPIRequest)
}

// ReleaseCainiaoCloudprintStdtemplatesGetAPIRequest 将 CainiaoCloudprintStdtemplatesGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintStdtemplatesGetAPIRequest(v *CainiaoCloudprintStdtemplatesGetAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintStdtemplatesGetAPIRequest.Put(v)
}
