package alihealthpw

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest 同步患者姓名至阿里健康 API请求
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
type AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest struct {
	model.Params
	// 入参
	_body *SynchroPatientNameDto
}

// NewAlibabaAlihealthPwSpecialSynchropatientnameRequest 初始化AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest对象
func NewAlibabaAlihealthPwSpecialSynchropatientnameRequest() *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest {
	return &AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.special.synchropatientname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) SetBody(_body *SynchroPatientNameDto) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) GetBody() *SynchroPatientNameDto {
	return r._body
}

var poolAlibabaAlihealthPwSpecialSynchropatientnameAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwSpecialSynchropatientnameRequest()
	},
}

// GetAlibabaAlihealthPwSpecialSynchropatientnameRequest 从 sync.Pool 获取 AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest
func GetAlibabaAlihealthPwSpecialSynchropatientnameAPIRequest() *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest {
	return poolAlibabaAlihealthPwSpecialSynchropatientnameAPIRequest.Get().(*AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest)
}

// ReleaseAlibabaAlihealthPwSpecialSynchropatientnameAPIRequest 将 AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwSpecialSynchropatientnameAPIRequest(v *AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwSpecialSynchropatientnameAPIRequest.Put(v)
}
