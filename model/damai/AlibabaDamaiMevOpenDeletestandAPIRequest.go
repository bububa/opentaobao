package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletestandAPIRequest
大麦换验平台-第三方对外开放-看台接口deleteStand API请求
alibaba.damai.mev.open.deletestand

deleteStand */
type AlibabaDamaiMevOpenDeletestandAPIRequest struct {
	model.Params
	// 入参deleteStandParam
	_deleteStandParam *StandIdOpenParam
}

// New
