package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripProjectModifyAPIRequest
变更项目 API请求
alitrip.btrip.project.modify

变更项目 */
type AlitripBtripProjectModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// New
