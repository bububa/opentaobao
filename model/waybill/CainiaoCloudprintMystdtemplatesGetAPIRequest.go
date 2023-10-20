package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintMystdtemplatesGetAPIRequest 获取用户使用的菜鸟电子面单模板信息 API请求
// cainiao.cloudprint.mystdtemplates.get
//
// 获取用户使用的菜鸟电子面单
type CainiaoCloudprintMystdtemplatesGetAPIRequest struct {
	model.Params
}

// NewCainiaoCloudprintMystdtemplatesGetRequest 初始化CainiaoCloudprintMystdtemplatesGetAPIRequest对象
func NewCainiaoCloudprintMystdtemplatesGetRequest() *CainiaoCloudprintMystdtemplatesGetAPIRequest {
	return &CainiaoCloudprintMystdtemplatesGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintMystdtemplatesGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintMystdtemplatesGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.mystdtemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintMystdtemplatesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintMystdtemplatesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolCainiaoCloudprintMystdtemplatesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintMystdtemplatesGetRequest()
	},
}

// GetCainiaoCloudprintMystdtemplatesGetRequest 从 sync.Pool 获取 CainiaoCloudprintMystdtemplatesGetAPIRequest
func GetCainiaoCloudprintMystdtemplatesGetAPIRequest() *CainiaoCloudprintMystdtemplatesGetAPIRequest {
	return poolCainiaoCloudprintMystdtemplatesGetAPIRequest.Get().(*CainiaoCloudprintMystdtemplatesGetAPIRequest)
}

// ReleaseCainiaoCloudprintMystdtemplatesGetAPIRequest 将 CainiaoCloudprintMystdtemplatesGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintMystdtemplatesGetAPIRequest(v *CainiaoCloudprintMystdtemplatesGetAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintMystdtemplatesGetAPIRequest.Put(v)
}
