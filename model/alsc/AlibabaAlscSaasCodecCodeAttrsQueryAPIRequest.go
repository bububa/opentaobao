package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest
码业务属性查询 API请求
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询 */
type AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest struct {
	model.Params
	// 请求入参
	_queryCodeRequest *QueryCodeBizAttrRequest
}

// New
