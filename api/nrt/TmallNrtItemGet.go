package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtitemget 家装新零售商品信息查询
// tmall.nrt.item.get
//
// 查询新零售商品信息
func Tmallnrtitemget(clt *core.SDKClient, req *nrt.TmallnrtitemgetAPIRequest, session string) (*nrt.TmallnrtitemgetAPIResponse, error) {
	var resp nrt.TmallnrtitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
