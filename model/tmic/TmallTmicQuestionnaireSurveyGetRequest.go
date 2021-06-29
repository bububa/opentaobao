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
    _hashCode   string
    // biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
    _biz   string
    // open_id
    _openUserId   string
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
func (r *TmallTmicQuestionnaireSurveyGetRequest) SetHashCode(_hashCode string) error {
    r._hashCode = _hashCode
    r.Set("hash_code", _hashCode)
    return nil
}

// HashCode Getter
func (r TmallTmicQuestionnaireSurveyGetRequest) GetHashCode() string {
    return r._hashCode
}
// Biz Setter
// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
func (r *TmallTmicQuestionnaireSurveyGetRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r TmallTmicQuestionnaireSurveyGetRequest) GetBiz() string {
    return r._biz
}
// OpenUserId Setter
// open_id
func (r *TmallTmicQuestionnaireSurveyGetRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TmallTmicQuestionnaireSurveyGetRequest) GetOpenUserId() string {
    return r._openUserId
}
