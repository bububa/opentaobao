package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardevaluateAPIRequest 服务商反馈鉴定结果 API请求
// tmall.servicecenter.workcard.evaluate
//
// 服务商反馈鉴定结果
type TmallservicecenterworkcardevaluateAPIRequest struct {
	model.Params
	// 鉴定结果图片列表
	_picUrlList []string
	// 鉴定不通过时的原因编码
	_failCode int64
	// 工单id
	_workcardId int64
	// 是否鉴定通过
	_passEvaluation bool
}

// NewTmallservicecenterworkcardevaluateRequest 初始化TmallservicecenterworkcardevaluateAPIRequest对象
func NewTmallservicecenterworkcardevaluateRequest() *TmallservicecenterworkcardevaluateAPIRequest {
	return &TmallservicecenterworkcardevaluateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardevaluateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardevaluateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardevaluateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicUrlList is PicUrlList Setter
// 鉴定结果图片列表
func (r *TmallservicecenterworkcardevaluateAPIRequest) SetPicUrlList(_picUrlList []string) error {
	r._picUrlList = _picUrlList
	r.Set("pic_url_list", _picUrlList)
	return nil
}

// GetPicUrlList PicUrlList Getter
func (r TmallservicecenterworkcardevaluateAPIRequest) GetPicUrlList() []string {
	return r._picUrlList
}

// SetFailCode is FailCode Setter
// 鉴定不通过时的原因编码
func (r *TmallservicecenterworkcardevaluateAPIRequest) SetFailCode(_failCode int64) error {
	r._failCode = _failCode
	r.Set("fail_code", _failCode)
	return nil
}

// GetFailCode FailCode Getter
func (r TmallservicecenterworkcardevaluateAPIRequest) GetFailCode() int64 {
	return r._failCode
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardevaluateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardevaluateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetPassEvaluation is PassEvaluation Setter
// 是否鉴定通过
func (r *TmallservicecenterworkcardevaluateAPIRequest) SetPassEvaluation(_passEvaluation bool) error {
	r._passEvaluation = _passEvaluation
	r.Set("pass_evaluation", _passEvaluation)
	return nil
}

// GetPassEvaluation PassEvaluation Getter
func (r TmallservicecenterworkcardevaluateAPIRequest) GetPassEvaluation() bool {
	return r._passEvaluation
}
