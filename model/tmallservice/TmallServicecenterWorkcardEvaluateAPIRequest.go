package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardEvaluateAPIRequest 服务商反馈鉴定结果 API请求
// tmall.servicecenter.workcard.evaluate
//
// 服务商反馈鉴定结果
type TmallServicecenterWorkcardEvaluateAPIRequest struct {
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

// NewTmallServicecenterWorkcardEvaluateRequest 初始化TmallServicecenterWorkcardEvaluateAPIRequest对象
func NewTmallServicecenterWorkcardEvaluateRequest() *TmallServicecenterWorkcardEvaluateAPIRequest {
	return &TmallServicecenterWorkcardEvaluateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicUrlList is PicUrlList Setter
// 鉴定结果图片列表
func (r *TmallServicecenterWorkcardEvaluateAPIRequest) SetPicUrlList(_picUrlList []string) error {
	r._picUrlList = _picUrlList
	r.Set("pic_url_list", _picUrlList)
	return nil
}

// GetPicUrlList PicUrlList Getter
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetPicUrlList() []string {
	return r._picUrlList
}

// SetFailCode is FailCode Setter
// 鉴定不通过时的原因编码
func (r *TmallServicecenterWorkcardEvaluateAPIRequest) SetFailCode(_failCode int64) error {
	r._failCode = _failCode
	r.Set("fail_code", _failCode)
	return nil
}

// GetFailCode FailCode Getter
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetFailCode() int64 {
	return r._failCode
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardEvaluateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetPassEvaluation is PassEvaluation Setter
// 是否鉴定通过
func (r *TmallServicecenterWorkcardEvaluateAPIRequest) SetPassEvaluation(_passEvaluation bool) error {
	r._passEvaluation = _passEvaluation
	r.Set("pass_evaluation", _passEvaluation)
	return nil
}

// GetPassEvaluation PassEvaluation Getter
func (r TmallServicecenterWorkcardEvaluateAPIRequest) GetPassEvaluation() bool {
	return r._passEvaluation
}
