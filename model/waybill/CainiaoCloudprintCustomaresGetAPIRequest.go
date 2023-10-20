package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintCustomaresGetAPIRequest 获取商家的自定义区模板信息 API请求
// cainiao.cloudprint.customares.get
//
// 供isv使用，获取商家的自定义区的模板信息
type CainiaoCloudprintCustomaresGetAPIRequest struct {
	model.Params
	// 用户使用的标准模板id
	_templateId int64
}

// NewCainiaoCloudprintCustomaresGetRequest 初始化CainiaoCloudprintCustomaresGetAPIRequest对象
func NewCainiaoCloudprintCustomaresGetRequest() *CainiaoCloudprintCustomaresGetAPIRequest {
	return &CainiaoCloudprintCustomaresGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintCustomaresGetAPIRequest) Reset() {
	r._templateId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCustomaresGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.customares.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCustomaresGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintCustomaresGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 用户使用的标准模板id
func (r *CainiaoCloudprintCustomaresGetAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r CainiaoCloudprintCustomaresGetAPIRequest) GetTemplateId() int64 {
	return r._templateId
}

var poolCainiaoCloudprintCustomaresGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintCustomaresGetRequest()
	},
}

// GetCainiaoCloudprintCustomaresGetRequest 从 sync.Pool 获取 CainiaoCloudprintCustomaresGetAPIRequest
func GetCainiaoCloudprintCustomaresGetAPIRequest() *CainiaoCloudprintCustomaresGetAPIRequest {
	return poolCainiaoCloudprintCustomaresGetAPIRequest.Get().(*CainiaoCloudprintCustomaresGetAPIRequest)
}

// ReleaseCainiaoCloudprintCustomaresGetAPIRequest 将 CainiaoCloudprintCustomaresGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintCustomaresGetAPIRequest(v *CainiaoCloudprintCustomaresGetAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintCustomaresGetAPIRequest.Put(v)
}
