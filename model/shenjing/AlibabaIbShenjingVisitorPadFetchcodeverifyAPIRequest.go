package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest
访客通过PAD提交访客码 API请求
alibaba.ib.shenjing.visitor.pad.fetchcodeverify

访客通过PAD提交访客码，录脸进入园区。 */
type AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest struct {
	model.Params
	// 访客码
	_visitorCode int64
	// 终端ID
	_termId string
}

// New
