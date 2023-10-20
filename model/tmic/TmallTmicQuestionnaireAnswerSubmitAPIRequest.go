package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltmicquestionnaireanswersubmitAPIRequest 提交问卷答案 API请求
// tmall.tmic.questionnaire.answer.submit
//
// 天猫新品创新中心对外开放问卷，提交问卷答案
type TmalltmicquestionnaireanswersubmitAPIRequest struct {
	model.Params
	// 用户填写的回答，类型为数组
	_userAnswerList []AnswerBo
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// openId
	_openUserId string
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
}

// NewTmalltmicquestionnaireanswersubmitRequest 初始化TmalltmicquestionnaireanswersubmitAPIRequest对象
func NewTmalltmicquestionnaireanswersubmitRequest() *TmalltmicquestionnaireanswersubmitAPIRequest {
	return &TmalltmicquestionnaireanswersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.answer.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserAnswerList is UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmalltmicquestionnaireanswersubmitAPIRequest) SetUserAnswerList(_userAnswerList []AnswerBo) error {
	r._userAnswerList = _userAnswerList
	r.Set("user_answer_list", _userAnswerList)
	return nil
}

// GetUserAnswerList UserAnswerList Getter
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetUserAnswerList() []AnswerBo {
	return r._userAnswerList
}

// SetHashCode is HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmalltmicquestionnaireanswersubmitAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// GetHashCode HashCode Getter
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetHashCode() string {
	return r._hashCode
}

// SetBiz is Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmalltmicquestionnaireanswersubmitAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetBiz() string {
	return r._biz
}

// SetOpenUserId is OpenUserId Setter
// openId
func (r *TmalltmicquestionnaireanswersubmitAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetRecordId is RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmalltmicquestionnaireanswersubmitAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// SetVersion is Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmalltmicquestionnaireanswersubmitAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmalltmicquestionnaireanswersubmitAPIRequest) GetVersion() int64 {
	return r._version
}
