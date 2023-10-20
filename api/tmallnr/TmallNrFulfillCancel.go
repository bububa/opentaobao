package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrfulfillcancel 取消上门揽件
// tmall.nr.fulfill.cancel
//
// 新零售门店业务取消上门揽件
func Tmallnrfulfillcancel(clt *core.SDKClient, req *tmallnr.TmallnrfulfillcancelAPIRequest, session string) (*tmallnr.TmallnrfulfillcancelAPIResponse, error) {
	var resp tmallnr.TmallnrfulfillcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
