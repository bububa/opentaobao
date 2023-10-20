package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrsellerstorerangeread 门店服务范围读取
// tmall.nr.seller.storerange.read
//
// 读取卖家所属门店的服务范围
func Tmallnrsellerstorerangeread(clt *core.SDKClient, req *tmallnr.TmallnrsellerstorerangereadAPIRequest, session string) (*tmallnr.TmallnrsellerstorerangereadAPIResponse, error) {
	var resp tmallnr.TmallnrsellerstorerangereadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
