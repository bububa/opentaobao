package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireOptionGetAPIRequest 获取单题选项 API请求
// tmall.tmic.questionnaire.option.get
//
// 根据具体题号，获取当前题目的选项列表
type TmallTmicQuestionnaireOptionGetAPIRequest struct {
	model.Params
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
	_questionCode string
	// 业务扩展参数
	_extraParameters string
	// openId
	_openUserId string
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
}

// NewTmallTmicQuestionnaireOptionGetRequest 初始化TmallTmicQuestionnaireOptionGetAPIRequest对象
func NewTmallTmicQuestionnaireOptionGetRequest() *TmallTmicQuestionnaireOptionGetAPIRequest {
	return &TmallTmicQuestionnaireOptionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.option.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHashCode is HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// GetHashCode HashCode Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetHashCode() string {
	return r._hashCode
}

// SetBiz is Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetBiz() string {
	return r._biz
}

// SetQuestionCode is QuestionCode Setter
// 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetQuestionCode(_questionCode string) error {
	r._questionCode = _questionCode
	r.Set("question_code", _questionCode)
	return nil
}

// GetQuestionCode QuestionCode Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetQuestionCode() string {
	return r._questionCode
}

// SetExtraParameters is ExtraParameters Setter
// 业务扩展参数
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetExtraParameters(_extraParameters string) error {
	r._extraParameters = _extraParameters
	r.Set("extra_parameters", _extraParameters)
	return nil
}

// GetExtraParameters ExtraParameters Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetExtraParameters() string {
	return r._extraParameters
}

// SetOpenUserId is OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetVersion is Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetVersion() int64 {
	return r._version
}

// SetRecordId is RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TmallTmicQuestionnaireOptionGetAPIRequest) GetRecordId() int64 {
	return r._recordId
}
