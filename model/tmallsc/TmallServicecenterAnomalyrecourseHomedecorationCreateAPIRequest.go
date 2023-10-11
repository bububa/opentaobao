package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest 天猫服务平台服务商代商家发起投诉单 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.create
//
// 天猫服务平台服务商代商家发起投诉单
type TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest struct {
	model.Params
	// 问题code
	_questionCode string
	// 投诉描述
	_remark string
	// 投诉图片url
	_images string
	// 工单ID
	_workcardId int64
	// 是否有消费者投诉风险0：无；1：有
	_publicOpinion int64
	// 投诉份数
	_complainCount int64
	// 申请赔付金额
	_applyCompensationAmount int64
}

// NewTmallServicecenterAnomalyrecourseHomedecorationCreateRequest 初始化TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest对象
func NewTmallServicecenterAnomalyrecourseHomedecorationCreateRequest() *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest {
	return &TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuestionCode is QuestionCode Setter
// 问题code
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetQuestionCode(_questionCode string) error {
	r._questionCode = _questionCode
	r.Set("question_code", _questionCode)
	return nil
}

// GetQuestionCode QuestionCode Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetQuestionCode() string {
	return r._questionCode
}

// SetRemark is Remark Setter
// 投诉描述
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetRemark() string {
	return r._remark
}

// SetImages is Images Setter
// 投诉图片url
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetImages(_images string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetImages() string {
	return r._images
}

// SetWorkcardId is WorkcardId Setter
// 工单ID
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetPublicOpinion is PublicOpinion Setter
// 是否有消费者投诉风险0：无；1：有
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetPublicOpinion(_publicOpinion int64) error {
	r._publicOpinion = _publicOpinion
	r.Set("public_opinion", _publicOpinion)
	return nil
}

// GetPublicOpinion PublicOpinion Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetPublicOpinion() int64 {
	return r._publicOpinion
}

// SetComplainCount is ComplainCount Setter
// 投诉份数
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetComplainCount(_complainCount int64) error {
	r._complainCount = _complainCount
	r.Set("complain_count", _complainCount)
	return nil
}

// GetComplainCount ComplainCount Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetComplainCount() int64 {
	return r._complainCount
}

// SetApplyCompensationAmount is ApplyCompensationAmount Setter
// 申请赔付金额
func (r *TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) SetApplyCompensationAmount(_applyCompensationAmount int64) error {
	r._applyCompensationAmount = _applyCompensationAmount
	r.Set("apply_compensation_amount", _applyCompensationAmount)
	return nil
}

// GetApplyCompensationAmount ApplyCompensationAmount Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest) GetApplyCompensationAmount() int64 {
	return r._applyCompensationAmount
}
