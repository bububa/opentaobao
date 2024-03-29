package tmic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireAnswerPushAPIRequest 提交单题答案 API请求
// tmall.tmic.questionnaire.answer.push
//
// 问卷单题回答的提交
type TmallTmicQuestionnaireAnswerPushAPIRequest struct {
	model.Params
	// 用户填写的回答，类型为数组
	_userAnswerList []AnswerBo
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// 开发平台userId
	_openUserId string
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
}

// NewTmallTmicQuestionnaireAnswerPushRequest 初始化TmallTmicQuestionnaireAnswerPushAPIRequest对象
func NewTmallTmicQuestionnaireAnswerPushRequest() *TmallTmicQuestionnaireAnswerPushAPIRequest {
	return &TmallTmicQuestionnaireAnswerPushAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) Reset() {
	r._userAnswerList = r._userAnswerList[:0]
	r._hashCode = ""
	r._biz = ""
	r._openUserId = ""
	r._recordId = 0
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.answer.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserAnswerList is UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) SetUserAnswerList(_userAnswerList []AnswerBo) error {
	r._userAnswerList = _userAnswerList
	r.Set("user_answer_list", _userAnswerList)
	return nil
}

// GetUserAnswerList UserAnswerList Getter
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetUserAnswerList() []AnswerBo {
	return r._userAnswerList
}

// SetHashCode is HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// GetHashCode HashCode Getter
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetHashCode() string {
	return r._hashCode
}

// SetBiz is Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetBiz() string {
	return r._biz
}

// SetOpenUserId is OpenUserId Setter
// 开发平台userId
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetOpenUserId() string {
	return r._openUserId
}

// SetRecordId is RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) SetRecordId(_recordId int64) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetRecordId() int64 {
	return r._recordId
}

// SetVersion is Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerPushAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmallTmicQuestionnaireAnswerPushAPIRequest) GetVersion() int64 {
	return r._version
}

var poolTmallTmicQuestionnaireAnswerPushAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTmicQuestionnaireAnswerPushRequest()
	},
}

// GetTmallTmicQuestionnaireAnswerPushRequest 从 sync.Pool 获取 TmallTmicQuestionnaireAnswerPushAPIRequest
func GetTmallTmicQuestionnaireAnswerPushAPIRequest() *TmallTmicQuestionnaireAnswerPushAPIRequest {
	return poolTmallTmicQuestionnaireAnswerPushAPIRequest.Get().(*TmallTmicQuestionnaireAnswerPushAPIRequest)
}

// ReleaseTmallTmicQuestionnaireAnswerPushAPIRequest 将 TmallTmicQuestionnaireAnswerPushAPIRequest 放入 sync.Pool
func ReleaseTmallTmicQuestionnaireAnswerPushAPIRequest(v *TmallTmicQuestionnaireAnswerPushAPIRequest) {
	v.Reset()
	poolTmallTmicQuestionnaireAnswerPushAPIRequest.Put(v)
}
