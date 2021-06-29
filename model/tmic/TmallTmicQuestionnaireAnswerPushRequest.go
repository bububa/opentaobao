package tmic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交单题答案 APIRequest
tmall.tmic.questionnaire.answer.push

问卷单题回答的提交
*/
type TmallTmicQuestionnaireAnswerPushRequest struct {
    model.Params

    // 问卷填答id，从问卷信息接口的应答中获取
    recordId   int64 

    // 问卷唯一编码，从问卷信息接口应答中获取
    hashCode   string 

    // 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
    biz   string 

    // 问卷版本号，从问卷信息接口的应答中获取
    version   int64 

    // 用户填写的回答，类型为数组
    userAnswerList   []AnswerBo 

    // 开发平台userId
    openUserId   string 

}

func NewTmallTmicQuestionnaireAnswerPushRequest() *TmallTmicQuestionnaireAnswerPushRequest{
    return &TmallTmicQuestionnaireAnswerPushRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetApiMethodName() string {
    return "tmall.tmic.questionnaire.answer.push"
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTmicQuestionnaireAnswerPushRequest) SetRecordId(recordId int64) error {
    r.recordId = recordId
    r.Set("record_id", recordId)
    return nil
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetRecordId() int64 {
    return r.recordId
}

func (r *TmallTmicQuestionnaireAnswerPushRequest) SetHashCode(hashCode string) error {
    r.hashCode = hashCode
    r.Set("hash_code", hashCode)
    return nil
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetHashCode() string {
    return r.hashCode
}

func (r *TmallTmicQuestionnaireAnswerPushRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetBiz() string {
    return r.biz
}

func (r *TmallTmicQuestionnaireAnswerPushRequest) SetVersion(version int64) error {
    r.version = version
    r.Set("version", version)
    return nil
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetVersion() int64 {
    return r.version
}

func (r *TmallTmicQuestionnaireAnswerPushRequest) SetUserAnswerList(userAnswerList []AnswerBo) error {
    r.userAnswerList = userAnswerList
    r.Set("user_answer_list", userAnswerList)
    return nil
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetUserAnswerList() []AnswerBo {
    return r.userAnswerList
}

func (r *TmallTmicQuestionnaireAnswerPushRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

func (r TmallTmicQuestionnaireAnswerPushRequest) GetOpenUserId() string {
    return r.openUserId
}

