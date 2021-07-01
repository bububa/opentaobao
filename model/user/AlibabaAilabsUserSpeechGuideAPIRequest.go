package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsUserSpeechGuideAPIRequest
引导语推荐接口 API请求
alibaba.ailabs.user.speech.guide

根据用户的语音query，返回相应的引导语推荐 */
type AlibabaAilabsUserSpeechGuideAPIRequest struct {
	model.Params
	// 用户query
	_query string
}

// New
