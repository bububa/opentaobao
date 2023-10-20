package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrsellerstorerangesync 同步商户中心服务范围
// tmall.nr.seller.storerange.sync
//
// 同步商户中心服务范围
func Tmallnrsellerstorerangesync(clt *core.SDKClient, req *tmallnr.TmallnrsellerstorerangesyncAPIRequest, session string) (*tmallnr.TmallnrsellerstorerangesyncAPIResponse, error) {
	var resp tmallnr.TmallnrsellerstorerangesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
