package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabUserOpenUidGetAPIRequest
access token 获取精灵用户 id API请求
alibaba.ailab.user.open.uid.get

access token 获取精灵用户 id */
type AlibabaAilabUserOpenUidGetAPIRequest struct {
	model.Params
	// access token
	_skillAccessToken string
	// skill id
	_skillId int64
}

// New
