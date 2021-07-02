package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTmicQuestionnaireSurveyGetAPIRequest 天猫新品创新中心问卷数据获取 API请求
// tmall.tmic.questionnaire.survey.get
//
// 天猫新品创新中心问卷数据获取
type TmallTmicQuestionnaireSurveyGetAPIRequest struct {
	model.Params
	// 问卷hashCode，问卷对外唯一编码
	_hashCode string
	// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
	_biz string
	// open_id
	_openUserId string
}

// NewTmallTmicQuestionnaireSurveyGetRequest 初始化TmallTmicQuestionnaireSurveyGetAPIRequest对象
func NewTmallTmicQuestionnaireSurveyGetRequest() *TmallTmicQuestionnaireSurveyGetAPIRequest {
	return &TmallTmicQuestionnaireSurveyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmicQuestionnaireSurveyGetAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.survey.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmicQuestionnaireSurveyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is HashCode Setter
// 问卷hashCode，问卷对外唯一编码
func (r *TmallTmicQuestionnaireSurveyGetAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// Get HashCode Getter
func (r TmallTmicQuestionnaireSurveyGetAPIRequest) GetHashCode() string {
	return r._hashCode
}

// Set is Biz Setter
// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
func (r *TmallTmicQuestionnaireSurveyGetAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// Get Biz Getter
func (r TmallTmicQuestionnaireSurveyGetAPIRequest) GetBiz() string {
	return r._biz
}

// Set is OpenUserId Setter
// open_id
func (r *TmallTmicQuestionnaireSurveyGetAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// Get OpenUserId Getter
func (r TmallTmicQuestionnaireSurveyGetAPIRequest) GetOpenUserId() string {
	return r._openUserId
}
