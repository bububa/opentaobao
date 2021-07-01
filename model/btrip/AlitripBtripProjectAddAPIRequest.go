package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripProjectAddAPIRequest
添加项目 API请求
alitrip.btrip.project.add

添加项目 */
type AlitripBtripProjectAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenProjectRq
}

// New
