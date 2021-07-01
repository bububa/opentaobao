package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest
获取药品扫码落地页vivo API请求
alibaba.alihealth.tracecodesearch.getshowurl.vivo

获取药品扫码落地页vivo */
type AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 来源
	_channel string
}

// New
