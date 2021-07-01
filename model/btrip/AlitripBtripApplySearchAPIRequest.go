package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripApplySearchAPIRequest
搜索审批单 API请求
alitrip.btrip.apply.search

外部企业调用获取本企业审批单列表数据 */
type AlitripBtripApplySearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSearchRq
}

// New
