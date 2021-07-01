package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteFaceelementAPIRequest
大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement API请求
alibaba.damai.mev.open.delete.faceelement

deleteFaceElement */
type AlibabaDamaiMevOpenDeleteFaceelementAPIRequest struct {
	model.Params
	// 入参deleteFaceElementParam
	_deleteFaceElementParam *FaceElementIdOpenParam
}

// New
