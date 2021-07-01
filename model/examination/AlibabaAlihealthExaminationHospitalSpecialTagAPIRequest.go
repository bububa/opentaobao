package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest
体检机构获取特色服务标签 API请求
alibaba.alihealth.examination.hospital.special.tag

体检机构获取特色服务标签列表 */
type AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest struct {
	model.Params
}

// NewAlibabaAlihealthExaminationHospitalSpecialTagRequest 初始化AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest对象
func NewAlibabaAlihealthExaminationHospitalSpecialTagRequest() *AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest {
	return &AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.hospital.special.tag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
