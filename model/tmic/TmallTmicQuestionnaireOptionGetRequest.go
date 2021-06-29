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
    hashCode   string
    // 问卷版本号，从问卷信息接口的应答中获取
    version   int64
    // 问卷填答id，从问卷信息接口的应答中获取
    recordId   int64
    // 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
    biz   string
    // 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
    questionCode   string
    // 业务扩展参数
    extraParameters   string
    // openId
    openUserId   string
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
func (r *TmallTmicQuestionnaireOptionGetRequest) SetHashCode(hashCode string) error {
    r.hashCode = hashCode
    r.Set("hash_code", hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetHashCode() string {
    return r.hashCode
}
// Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetVersion() int64 {
    return r.version
}
// RecordId Setter
// 问卷填答id，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetRecordId(recordId int64) error {
    r.recordId = recordId
    r.Set("record_id", recordId)
    return nil
}

// RecordId Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetRecordId() int64 {
    return r.recordId
}
// Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireOptionGetRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetBiz() string {
    return r.biz
}
// QuestionCode Setter
// 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireOptionGetRequest) SetQuestionCode(questionCode string) error {
    r.questionCode = questionCode
    r.Set("question_code", questionCode)
    return nil
}

// QuestionCode Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetQuestionCode() string {
    return r.questionCode
}
// ExtraParameters Setter
// 业务扩展参数
func (r *TmallTmicQuestionnaireOptionGetRequest) SetExtraParameters(extraParameters string) error {
    r.extraParameters = extraParameters
    r.Set("extra_parameters", extraParameters)
    return nil
}

// ExtraParameters Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetExtraParameters() string {
    return r.extraParameters
}
// OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireOptionGetRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireOptionGetRequest) GetOpenUserId() string {
    return r.openUserId
}
