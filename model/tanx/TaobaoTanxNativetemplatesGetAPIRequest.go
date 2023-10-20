package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxnativetemplatesgetAPIRequest 批量获取本地模板信息 API请求
// taobao.tanx.nativetemplates.get
//
// 根据传入的本地模板ID批量返回本地模板
type TaobaotanxnativetemplatesgetAPIRequest struct {
	model.Params
	// 本地模板ID列表
	_templateIds []int64
	// dsp对应的tanx的token
	_token string
	// dsp在tanx的memberid
	_memberId int64
	// 1970年到现在的毫秒
	_signTime int64
}

// NewTaobaotanxnativetemplatesgetRequest 初始化TaobaotanxnativetemplatesgetAPIRequest对象
func NewTaobaotanxnativetemplatesgetRequest() *TaobaotanxnativetemplatesgetAPIRequest {
	return &TaobaotanxnativetemplatesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxnativetemplatesgetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.nativetemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxnativetemplatesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxnativetemplatesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateIds is TemplateIds Setter
// 本地模板ID列表
func (r *TaobaotanxnativetemplatesgetAPIRequest) SetTemplateIds(_templateIds []int64) error {
	r._templateIds = _templateIds
	r.Set("template_ids", _templateIds)
	return nil
}

// GetTemplateIds TemplateIds Getter
func (r TaobaotanxnativetemplatesgetAPIRequest) GetTemplateIds() []int64 {
	return r._templateIds
}

// SetToken is Token Setter
// dsp对应的tanx的token
func (r *TaobaotanxnativetemplatesgetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxnativetemplatesgetAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// dsp在tanx的memberid
func (r *TaobaotanxnativetemplatesgetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxnativetemplatesgetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 1970年到现在的毫秒
func (r *TaobaotanxnativetemplatesgetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxnativetemplatesgetAPIRequest) GetSignTime() int64 {
	return r._signTime
}
