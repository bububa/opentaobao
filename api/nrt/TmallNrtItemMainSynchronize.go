package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtItemMainSynchronize 家装新零售主商品同步至阿里
// tmall.nrt.item.main.synchronize
//
// 同步卖场存量线下商品到阿里
func TmallNrtItemMainSynchronize(clt *core.SDKClient, req *nrt.TmallNrtItemMainSynchronizeAPIRequest, resp *nrt.TmallNrtItemMainSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
