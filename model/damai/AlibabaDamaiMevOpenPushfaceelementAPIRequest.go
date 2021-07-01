package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushfaceelementAPIRequest
大麦换验平台-第三方对外开放-票面元素接口pushFaceElement API请求
alibaba.damai.mev.open.pushfaceelement

pushFaceElement */
type AlibabaDamaiMevOpenPushfaceelementAPIRequest struct {
	model.Params
	// 入参pushFaceElementParamList
	_pushFaceElementParamList []ThirdFaceElementPushOpenParam
}

// New
