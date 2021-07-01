package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushprojectAPIRequest
大麦换验平台-第三方对外开放-项目接口pushProject API请求
alibaba.damai.mev.open.pushproject

pushProject */
type AlibabaDamaiMevOpenPushprojectAPIRequest struct {
	model.Params
	// 入参pushProjectParam
	_pushProjectParam *ThirdProjectPushOpenParam
}

// New
