package tanx

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxNativetemplatesGetAPIRequest 批量获取本地模板信息 API请求
// taobao.tanx.nativetemplates.get
//
// 根据传入的本地模板ID批量返回本地模板
type TaobaoTanxNativetemplatesGetAPIRequest struct {
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

// NewTaobaoTanxNativetemplatesGetRequest 初始化TaobaoTanxNativetemplatesGetAPIRequest对象
func NewTaobaoTanxNativetemplatesGetRequest() *TaobaoTanxNativetemplatesGetAPIRequest {
	return &TaobaoTanxNativetemplatesGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTanxNativetemplatesGetAPIRequest) Reset() {
	r._templateIds = r._templateIds[:0]
	r._token = ""
	r._memberId = 0
	r._signTime = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.nativetemplates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTanxNativetemplatesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTanxNativetemplatesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTanxNativetemplatesGetRequest()
	},
}

// GetTaobaoTanxNativetemplatesGetRequest 从 sync.Pool 获取 TaobaoTanxNativetemplatesGetAPIRequest
func GetTaobaoTanxNativetemplatesGetAPIRequest() *TaobaoTanxNativetemplatesGetAPIRequest {
	return poolTaobaoTanxNativetemplatesGetAPIRequest.Get().(*TaobaoTanxNativetemplatesGetAPIRequest)
}

// ReleaseTaobaoTanxNativetemplatesGetAPIRequest 将 TaobaoTanxNativetemplatesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTanxNativetemplatesGetAPIRequest(v *TaobaoTanxNativetemplatesGetAPIRequest) {
	v.Reset()
	poolTaobaoTanxNativetemplatesGetAPIRequest.Put(v)
}
