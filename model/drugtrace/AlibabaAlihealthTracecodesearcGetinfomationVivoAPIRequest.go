package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest
获取vivo banner API请求
alibaba.alihealth.tracecodesearc.getinfomation.vivo

获取vivo banner  url */
type AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest struct {
	model.Params
	// 渠道
	_channel string
}

// New
