package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationhospitalspecialtagAPIRequest 体检机构获取特色服务标签 API请求
// alibaba.alihealth.examination.hospital.special.tag
//
// 体检机构获取特色服务标签列表
type AlibabaalihealthexaminationhospitalspecialtagAPIRequest struct {
	model.Params
}

// NewAlibabaalihealthexaminationhospitalspecialtagRequest 初始化AlibabaalihealthexaminationhospitalspecialtagAPIRequest对象
func NewAlibabaalihealthexaminationhospitalspecialtagRequest() *AlibabaalihealthexaminationhospitalspecialtagAPIRequest {
	return &AlibabaalihealthexaminationhospitalspecialtagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationhospitalspecialtagAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.hospital.special.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationhospitalspecialtagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationhospitalspecialtagAPIRequest) GetRawParams() model.Params {
	return r.Params
}
