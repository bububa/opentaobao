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
    recordId   int64
    // 用户填写的回答，类型为数组
    userAnswerList   []AnswerBo
    // 问卷唯一编码，从问卷信息接口应答中获取
    hashCode   string
    // 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
    biz   string
    // 问卷版本号，从问卷信息接口的应答中获取
    version   int64
    // openId
    openUserId   string
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
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetRecordId(recordId int64) error {
    r.recordId = recordId
    r.Set("record_id", recordId)
    return nil
}

// RecordId Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetRecordId() int64 {
    return r.recordId
}
// UserAnswerList Setter
// 用户填写的回答，类型为数组
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetUserAnswerList(userAnswerList []AnswerBo) error {
    r.userAnswerList = userAnswerList
    r.Set("user_answer_list", userAnswerList)
    return nil
}

// UserAnswerList Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetUserAnswerList() []AnswerBo {
    return r.userAnswerList
}
// HashCode Setter
// 问卷唯一编码，从问卷信息接口应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetHashCode(hashCode string) error {
    r.hashCode = hashCode
    r.Set("hash_code", hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetHashCode() string {
    return r.hashCode
}
// Biz Setter
// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetBiz() string {
    return r.biz
}
// Version Setter
// 问卷版本号，从问卷信息接口的应答中获取
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

// Version Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetVersion() int64 {
    return r.version
}
// OpenUserId Setter
// openId
func (r *TmallTmicQuestionnaireAnswerSubmitRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireAnswerSubmitRequest) GetOpenUserId() string {
    return r.openUserId
}
