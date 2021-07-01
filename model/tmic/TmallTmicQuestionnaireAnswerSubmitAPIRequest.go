package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTmicQuestionnaireAnswerSubmitAPIRequest
提交问卷答案 API请求
tmall.tmic.questionnaire.answer.submit

天猫新品创新中心对外开放问卷，提交问卷答案 */
type TmallTmicQuestionnaireAnswerSubmitAPIRequest struct {
	model.Params
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
	// 用户填写的回答，类型为数组
	_userAnswerList []AnswerBo
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
	// openId
	_openUserId string
}

// New
