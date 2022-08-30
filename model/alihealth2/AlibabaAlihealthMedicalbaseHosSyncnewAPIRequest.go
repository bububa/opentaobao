package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest 直连医院上传接口 API请求
// alibaba.alihealth.medicalbase.hos.syncnew
//
// 直连医院上传接口
type AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest struct {
	model.Params
	// 212
	_bizContent string
}

// NewAlibabaAlihealthMedicalbaseHosSyncnewRequest 初始化AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest对象
func NewAlibabaAlihealthMedicalbaseHosSyncnewRequest() *AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest {
	return &AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.hos.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizContent is BizContent Setter
// 212
func (r *AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) SetBizContent(_bizContent string) error {
	r._bizContent = _bizContent
	r.Set("biz_content", _bizContent)
	return nil
}

// GetBizContent BizContent Getter
func (r AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) GetBizContent() string {
	return r._bizContent
}
