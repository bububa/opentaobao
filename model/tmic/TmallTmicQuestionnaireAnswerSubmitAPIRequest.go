package tmic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireAnswerSubmitAPIRequest 提交问卷答案 API请求
// tmall.tmic.questionnaire.answer.submit
//
// 天猫新品创新中心对外开放问卷，提交问卷答案
type TmallTmicQuestionnaireAnswerSubmitAPIRequest struct {
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

// NewTmallTmicQuestionnaireAnswerSubmitRequest 初始化TmallTmicQuestionnaireAnswerSubmitAPIRequest对象
func NewTmallTmicQuestionnaireAnswerSubmitRequest() *TmallTmicQuestionnaireAnswerSubmitAPIRequest {
	return &TmallTmicQuestionnaireAnswerSubmitAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) Reset() {
	r._userAnswerList = r._userAnswerList[:0]
	r._hashCode = ""
	r._biz = ""
	r._openUserId = ""
	r._recordId = 0
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.answer.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserAnswerList is UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetUserAnswerList(_userAnswerList []AnswerBo) error {
	r._userAnswerList = _userAnswerList
	r.Set("user_answer_list", _userAnswerList)
	return nil
}

// GetUserAnswerList UserAnswerList Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetUserAnswerList() []AnswerBo {
	return r._userAnswerList
}

// SetHashCode is HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// GetHashCode HashCode Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetHashCode() string {
	return r._hashCode
}

// SetBiz is Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetBiz() string {
	return r._biz
}

// SetOpenUserId is OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetRecordId is RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// SetVersion is Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmallTmicQuestionnaireAnswerSubmitAPIRequest) GetVersion() int64 {
	return r._version
}

var poolTmallTmicQuestionnaireAnswerSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTmicQuestionnaireAnswerSubmitRequest()
	},
}

// GetTmallTmicQuestionnaireAnswerSubmitRequest 从 sync.Pool 获取 TmallTmicQuestionnaireAnswerSubmitAPIRequest
func GetTmallTmicQuestionnaireAnswerSubmitAPIRequest() *TmallTmicQuestionnaireAnswerSubmitAPIRequest {
	return poolTmallTmicQuestionnaireAnswerSubmitAPIRequest.Get().(*TmallTmicQuestionnaireAnswerSubmitAPIRequest)
}

// ReleaseTmallTmicQuestionnaireAnswerSubmitAPIRequest 将 TmallTmicQuestionnaireAnswerSubmitAPIRequest 放入 sync.Pool
func ReleaseTmallTmicQuestionnaireAnswerSubmitAPIRequest(v *TmallTmicQuestionnaireAnswerSubmitAPIRequest) {
	v.Reset()
	poolTmallTmicQuestionnaireAnswerSubmitAPIRequest.Put(v)
}
