package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireAnswerSubmitAPIRequest 提交问卷答案 API请求
// tmall.tmic.questionnaire.answer.submit
//
// 天猫新品创新中心对外开放问卷，提交问卷答案
type TmallTmicQuestionnaireAnswerSubmitAPIRequest struct {
	model.Params
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
	// 用户填写的回答，类型为数组
	_userAnswerList []AnswerBo
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
	// openId
	_openUserId string
}

// NewTmallTmicQuestionnaireAnswerSubmitRequest 初始化TmallTmicQuestionnaireAnswerSubmitAPIRequest对象
func NewTmallTmicQuestionnaireAnswerSubmitRequest() *TmallTmicQuestionnaireAnswerSubmitAPIRequest {
	return &TmallTmicQuestionnaireAnswerSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.answer.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// Get RecordId Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// Set is UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetUserAnswerList(_userAnswerList []AnswerBo) error {
	r._userAnswerList = _userAnswerList
	r.Set("user_answer_list", _userAnswerList)
	return nil
}

// Get UserAnswerList Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetUserAnswerList() []AnswerBo {
	return r._userAnswerList
}

// Set is HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// Get HashCode Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetHashCode() string {
	return r._hashCode
}

// Set is Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// Get Biz Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetBiz() string {
	return r._biz
}

// Set is Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetVersion() int64 {
	return r._version
}

// Set is OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// Get OpenUserId Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetOpenUserId() string {
	return r._openUserId
}
