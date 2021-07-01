package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripProjectDeleteAPIRequest
删除项目 API请求
alitrip.btrip.project.delete

删除项目 */
type AlitripBtripProjectDeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// New
