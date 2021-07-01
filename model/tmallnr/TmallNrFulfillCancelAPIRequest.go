package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillCancelAPIRequest
取消上门揽件 API请求
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件 */
type TmallNrFulfillCancelAPIRequest struct {
	model.Params
	// 入参
	_req *NrCancelFulfillReqDto
}

// New
