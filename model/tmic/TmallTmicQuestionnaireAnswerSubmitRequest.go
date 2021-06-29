package tmic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交问卷答案 API请求
tmall.tmic.questionnaire.answer.submit

天猫新品创新中心对外开放问卷，提交问卷答案
*/
type TmallTmicQuestionnaireAnswerSubmitRequest struct {
    model.Params
    // 问卷填答id，从问卷信息接口的应答中获取
    _recordId   int64
    // 用户填写的回答，类型为数组
    _userAnswerList   []AnswerBo
    // 问卷唯一编码，从问卷信息接口应答中获取
    _hashCode   string
    // 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
    _biz   string
    // 问卷版本号，从问卷信息接口的应答中获取
    _version   int64
    // openId
    _openUserId   string
}

// 初始化TmallTmicQuestionnaireAnswerSubmitRequest对象
func NewTmallTmicQuestionnaireAnswerSubmitRequest() *TmallTmicQuestionnaireAnswerSubmitRequest{
    return &TmallTmicQuestionnaireAnswerSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetApiMethodName() string {
    return "tmall.tmic.questionnaire.answer.submit"
}

// IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetRecordId(_recordId int64) error {
    r._recordId = _recordId
    r.Set("record_id", _recordId)
    return nil
}

// RecordId Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetRecordId() int64 {
    return r._recordId
}
// UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetUserAnswerList(_userAnswerList []AnswerBo) error {
    r._userAnswerList = _userAnswerList
    r.Set("user_answer_list", _userAnswerList)
    return nil
}

// UserAnswerList Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetUserAnswerList() []AnswerBo {
    return r._userAnswerList
}
// HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetHashCode(_hashCode string) error {
    r._hashCode = _hashCode
    r.Set("hash_code", _hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetHashCode() string {
    return r._hashCode
}
// Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetBiz() string {
    return r._biz
}
// Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetVersion() int64 {
    return r._version
}
// OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetOpenUserId() string {
    return r._openUserId
}
