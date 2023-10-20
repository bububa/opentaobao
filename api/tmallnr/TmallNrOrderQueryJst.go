package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrorderqueryjst 获取同城配送业务单笔订单
// tmall.nr.order.query.jst
//
// 同城配送业务获取单笔订单
func Tmallnrorderqueryjst(clt *core.SDKClient, req *tmallnr.TmallnrorderqueryjstAPIRequest, session string) (*tmallnr.TmallnrorderqueryjstAPIResponse, error) {
	var resp tmallnr.TmallnrorderqueryjstAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
