package alihealth2

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) Reset() {
	r._bizContent = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.hos.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthMedicalbaseHosSyncnewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthMedicalbaseHosSyncnewRequest()
	},
}

// GetAlibabaAlihealthMedicalbaseHosSyncnewRequest 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest
func GetAlibabaAlihealthMedicalbaseHosSyncnewAPIRequest() *AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest {
	return poolAlibabaAlihealthMedicalbaseHosSyncnewAPIRequest.Get().(*AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest)
}

// ReleaseAlibabaAlihealthMedicalbaseHosSyncnewAPIRequest 将 AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseHosSyncnewAPIRequest(v *AlibabaAlihealthMedicalbaseHosSyncnewAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseHosSyncnewAPIRequest.Put(v)
}
