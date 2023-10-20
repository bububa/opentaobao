package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrtnewcouponsend 券发放接口
// tmall.nrt.newcoupon.send
//
// 券发放接口
func Tmallnrtnewcouponsend(clt *core.SDKClient, req *tmallnr.TmallnrtnewcouponsendAPIRequest, session string) (*tmallnr.TmallnrtnewcouponsendAPIResponse, error) {
	var resp tmallnr.TmallnrtnewcouponsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
