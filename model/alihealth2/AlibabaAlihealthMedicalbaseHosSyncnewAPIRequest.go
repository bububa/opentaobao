package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasehossyncnewAPIRequest 直连医院上传接口 API请求
// alibaba.alihealth.medicalbase.hos.syncnew
//
// 直连医院上传接口
type AlibabaalihealthmedicalbasehossyncnewAPIRequest struct {
	model.Params
	// 212
	_bizContent string
}

// NewAlibabaalihealthmedicalbasehossyncnewRequest 初始化AlibabaalihealthmedicalbasehossyncnewAPIRequest对象
func NewAlibabaalihealthmedicalbasehossyncnewRequest() *AlibabaalihealthmedicalbasehossyncnewAPIRequest {
	return &AlibabaalihealthmedicalbasehossyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbasehossyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.hos.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbasehossyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbasehossyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizContent is BizContent Setter
// 212
func (r *AlibabaalihealthmedicalbasehossyncnewAPIRequest) SetBizContent(_bizContent string) error {
	r._bizContent = _bizContent
	r.Set("biz_content", _bizContent)
	return nil
}

// GetBizContent BizContent Getter
func (r AlibabaalihealthmedicalbasehossyncnewAPIRequest) GetBizContent() string {
	return r._bizContent
}
