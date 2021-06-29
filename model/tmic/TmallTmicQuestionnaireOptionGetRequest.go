package tmic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单题选项 API请求
tmall.tmic.questionnaire.option.get

根据具体题号，获取当前题目的选项列表
*/
type TmallTmicQuestionnaireOptionGetRequest struct {
    model.Params
    // 问卷唯一编码，从问卷信息接口应答中获取
    _hashCode   string
    // 问卷版本号，从问卷信息接口的应答中获取
    _version   int64
    // 问卷填答id，从问卷信息接口的应答中获取
    _recordId   int64
    // 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
    _biz   string
    // 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
    _questionCode   string
    // 业务扩展参数
    _extraParameters   string
    // openId
    _openUserId   string
}

// 初始化TmallTmicQuestionnaireOptionGetRequest对象
func NewTmallTmicQuestionnaireOptionGetRequest() *TmallTmicQuestionnaireOptionGetRequest{
    return &TmallTmicQuestionnaireOptionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireOptionGetRequest) GetApiMethodName() string {
    return "tmall.tmic.questionnaire.option.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireOptionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetHashCode(_hashCode string) error {
    r._hashCode = _hashCode
    r.Set("hash_code", _hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetHashCode() string {
    return r._hashCode
}
// Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetVersion(_version int64) error {
    r._version = _version
    r.Set("version", _version)
    return nil
}

// Version Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetVersion() int64 {
    return r._version
}
// RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetRecordId(_recordId int64) error {
    r._recordId = _recordId
    r.Set("record_id", _recordId)
    return nil
}

// RecordId Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetRecordId() int64 {
    return r._recordId
}
// Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireOptionGetRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetBiz() string {
    return r._biz
}
// QuestionCode Setter
// 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetQuestionCode(_questionCode string) error {
    r._questionCode = _questionCode
    r.Set("question_code", _questionCode)
    return nil
}

// QuestionCode Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetQuestionCode() string {
    return r._questionCode
}
// ExtraParameters Setter
// 业务扩展参数
func (r *TmallTmicQuestionnaireOptionGetRequest) SetExtraParameters(_extraParameters string) error {
    r._extraParameters = _extraParameters
    r.Set("extra_parameters", _extraParameters)
    return nil
}

// ExtraParameters Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetExtraParameters() string {
    return r._extraParameters
}
// OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireOptionGetRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetOpenUserId() string {
    return r._openUserId
}
