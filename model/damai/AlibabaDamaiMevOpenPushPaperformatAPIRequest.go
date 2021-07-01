package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushPaperformatAPIRequest
大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat API请求
alibaba.damai.mev.open.push.paperformat

pushPaperFormat */
type AlibabaDamaiMevOpenPushPaperformatAPIRequest struct {
	model.Params
	// 入参pushPaperFormatParam
	_pushPaperFormatParam *ThirdPaperFormatPushOpenParam
}

// New
