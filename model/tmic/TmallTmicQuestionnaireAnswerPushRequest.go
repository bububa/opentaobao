package tmic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交单题答案 API请求
tmall.tmic.questionnaire.answer.push

问卷单题回答的提交
*/
type TmallTmicQuestionnaireAnswerPushRequest struct {
    model.Params
    // 问卷填答id，从问卷信息接口的应答中获取
    _recordId   int64
    // 问卷唯一编码，从问卷信息接口应答中获取
    _hashCode   string
    // 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
    _biz   string
    // 问卷版本号，从问卷信息接口的应答中获取
    _version   int64
    // 用户填写的回答，类型为数组
    _userAnswerList   []AnswerBo
    // 开发平台userId
    _openUserId   string
}

// 初始化TmallTmicQuestionnaireAnswerPushRequest对象
func NewTmallTmicQuestionnaireAnswerPushRequest() *TmallTmicQuestionnaireAnswerPushRequest{
    return &TmallTmicQuestionnaireAnswerPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireAnswerPushRequest) GetApiMethodName() string {
    return "tmall.tmic.questionnaire.answer.push"
}

// IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireAnswerPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerPushRequest) SetRecordId(_recordId int64) error {
    r._recordId = _recordId
    r.Set("record_id", _recordId)
    return nil
}

// RecordId Getter
func (r TmallTmicQuestionnaireAnswerPushRequest) GetRecordId() int64 {
    return r._recordId
}
// HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireAnswerPushRequest) SetHashCode(_hashCode string) error {
    r._hashCode = _hashCode
    r.Set("hash_code", _hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireAnswerPushRequest) GetHashCode() string {
    return r._hashCode
}
// Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireAnswerPushRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireAnswerPushRequest) GetBiz() string {
    return r._biz
}
// Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerPushRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r TmallTmicQuestionnaireAnswerPushRequest) GetVersion() int64 {
    return r._version
}
// UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmallTmicQuestionnaireAnswerPushRequest) SetUserAnswerList(_userAnswerList []AnswerBo) error {
    r._userAnswerList = _userAnswerList
    r.Set("user_answer_list", _userAnswerList)
    return nil
}

// UserAnswerList Getter
func (r TmallTmicQuestionnaireAnswerPushRequest) GetUserAnswerList() []AnswerBo {
    return r._userAnswerList
}
// OpenUserId Setter
// 开发平台userId
func (r *TmallTmicQuestionnaireAnswerPushRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireAnswerPushRequest) GetOpenUserId() string {
    return r._openUserId
}