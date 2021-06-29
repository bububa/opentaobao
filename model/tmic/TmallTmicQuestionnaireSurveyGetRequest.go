package tmic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫新品创新中心问卷数据获取 API请求
tmall.tmic.questionnaire.survey.get

天猫新品创新中心问卷数据获取
*/
type TmallTmicQuestionnaireSurveyGetRequest struct {
    model.Params
    // 问卷hashCode，问卷对外唯一编码
    hashCode   string
    // biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
    biz   string
    // open_id
    openUserId   string
}

// 初始化TmallTmicQuestionnaireSurveyGetRequest对象
func NewTmallTmicQuestionnaireSurveyGetRequest() *TmallTmicQuestionnaireSurveyGetRequest{
    return &TmallTmicQuestionnaireSurveyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireSurveyGetRequest) GetApiMethodName() string {
    return "tmall.tmic.questionnaire.survey.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireSurveyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HashCode Setter
// 问卷hashCode，问卷对外唯一编码
func (r *TmallTmicQuestionnaireSurveyGetRequest) SetHashCode(hashCode string) error {
    r.hashCode = hashCode
    r.Set("hash_code", hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireSurveyGetRequest) GetHashCode() string {
    return r.hashCode
}
// Biz Setter
// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
func (r *TmallTmicQuestionnaireSurveyGetRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireSurveyGetRequest) GetBiz() string {
    return r.biz
}
// OpenUserId Setter
// open_id
func (r *TmallTmicQuestionnaireSurveyGetRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireSurveyGetRequest) GetOpenUserId() string {
    return r.openUserId
}
