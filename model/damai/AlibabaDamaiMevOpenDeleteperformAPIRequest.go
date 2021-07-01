package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteperformAPIRequest
大麦换验平台-第三方对外开放-场次接口deletePerform API请求
alibaba.damai.mev.open.deleteperform

deletePerform */
type AlibabaDamaiMevOpenDeleteperformAPIRequest struct {
	model.Params
	// 入参deletePerformParam
	_deletePerformParam *PerformIdOpenParam
}

// New
