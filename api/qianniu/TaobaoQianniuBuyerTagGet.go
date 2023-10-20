package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniubuyertagget 判断买家是否有某些标
// taobao.qianniu.buyer.tag.get
//
// 判断某个买家是否有某些标
func Taobaoqianniubuyertagget(clt *core.SDKClient, req *qianniu.TaobaoqianniubuyertaggetAPIRequest, session string) (*qianniu.TaobaoqianniubuyertaggetAPIResponse, error) {
	var resp qianniu.TaobaoqianniubuyertaggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
