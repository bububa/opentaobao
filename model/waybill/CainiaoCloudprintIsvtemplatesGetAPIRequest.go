package waybill

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintIsvtemplatesGetAPIRequest) Reset() {
	r.Params.ToZero()
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

var poolCainiaoCloudprintIsvtemplatesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintIsvtemplatesGetRequest()
	},
}

// GetCainiaoCloudprintIsvtemplatesGetRequest 从 sync.Pool 获取 CainiaoCloudprintIsvtemplatesGetAPIRequest
func GetCainiaoCloudprintIsvtemplatesGetAPIRequest() *CainiaoCloudprintIsvtemplatesGetAPIRequest {
	return poolCainiaoCloudprintIsvtemplatesGetAPIRequest.Get().(*CainiaoCloudprintIsvtemplatesGetAPIRequest)
}

// ReleaseCainiaoCloudprintIsvtemplatesGetAPIRequest 将 CainiaoCloudprintIsvtemplatesGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintIsvtemplatesGetAPIRequest(v *CainiaoCloudprintIsvtemplatesGetAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintIsvtemplatesGetAPIRequest.Put(v)
}
