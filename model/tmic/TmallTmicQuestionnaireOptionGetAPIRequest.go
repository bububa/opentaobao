package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTmicQuestionnaireOptionGetAPIRequest
获取单题选项 API请求
tmall.tmic.questionnaire.option.get

根据具体题号，获取当前题目的选项列表 */
type TmallTmicQuestionnaireOptionGetAPIRequest struct {
	model.Params
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// 问题编码，问卷中的问题的唯一编码，从问卷信息接口的应答中获取
	_questionCode string
	// 业务扩展参数
	_extraParameters string
	// openId
	_openUserId string
}

// New
