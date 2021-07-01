package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTmicQuestionnaireAnswerPushAPIRequest
提交单题答案 API请求
tmall.tmic.questionnaire.answer.push

问卷单题回答的提交 */
type TmallTmicQuestionnaireAnswerPushAPIRequest struct {
	model.Params
	// 问卷填答id，从问卷信息接口的应答中获取
	_recordId int64
	// 问卷唯一编码，从问卷信息接口应答中获取
	_hashCode string
	// 业务参数，区分问卷分组投放，1024表示分组投放id，fav表示用户动作类型为收藏
	_biz string
	// 问卷版本号，从问卷信息接口的应答中获取
	_version int64
	// 用户填写的回答，类型为数组
	_userAnswerList []AnswerBo
	// 开发平台userId
	_openUserId string
}

// New
