package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtItemGet 家装新零售商品信息查询
// tmall.nrt.item.get
//
// 查询新零售商品信息
func TmallNrtItemGet(clt *core.SDKClient, req *nrt.TmallNrtItemGetAPIRequest, resp *nrt.TmallNrtItemGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
