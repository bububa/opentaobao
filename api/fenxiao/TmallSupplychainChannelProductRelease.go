package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductRelease 供应商铺货
// tmall.supplychain.channel.product.release
//
// 供应商渠道铺货接口
func TmallSupplychainChannelProductRelease(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductReleaseAPIRequest, resp *fenxiao.TmallSupplychainChannelProductReleaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
