package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxNativetemplatesGetAPIRequest 批量获取本地模板信息 API请求
// taobao.tanx.nativetemplates.get
//
// 根据传入的本地模板ID批量返回本地模板
type TaobaoTanxNativetemplatesGetAPIRequest struct {
	model.Params
	// dsp在tanx的memberid
	_memberId int64
	// dsp对应的tanx的token
	_token string
	// 1970年到现在的毫秒
	_signTime int64
	// 本地模板ID列表
	_templateIds []int64
}

// NewTaobaoTanxNativetemplatesGetRequest 初始化TaobaoTanxNativetemplatesGetAPIRequest对象
func NewTaobaoTanxNativetemplatesGetRequest() *TaobaoTanxNativetemplatesGetAPIRequest {
	return &TaobaoTanxNativetemplatesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.nativetemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMemberId is MemberId Setter
// dsp在tanx的memberid
func (r *TaobaoTanxNativetemplatesGetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetToken is Token Setter
// dsp对应的tanx的token
func (r *TaobaoTanxNativetemplatesGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetToken() string {
	return r._token
}

// SetSignTime is SignTime Setter
// 1970年到现在的毫秒
func (r *TaobaoTanxNativetemplatesGetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetTemplateIds is TemplateIds Setter
// 本地模板ID列表
func (r *TaobaoTanxNativetemplatesGetAPIRequest) SetTemplateIds(_templateIds []int64) error {
	r._templateIds = _templateIds
	r.Set("template_ids", _templateIds)
	return nil
}

// GetTemplateIds TemplateIds Getter
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetTemplateIds() []int64 {
	return r._templateIds
}
