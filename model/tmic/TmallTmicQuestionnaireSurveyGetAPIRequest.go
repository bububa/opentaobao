package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTmicQuestionnaireSurveyGetAPIRequest
天猫新品创新中心问卷数据获取 API请求
tmall.tmic.questionnaire.survey.get

天猫新品创新中心问卷数据获取 */
type TmallTmicQuestionnaireSurveyGetAPIRequest struct {
	model.Params
	// 问卷hashCode，问卷对外唯一编码
	_hashCode string
	// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
	_biz string
	// open_id
	_openUserId string
}

// New
