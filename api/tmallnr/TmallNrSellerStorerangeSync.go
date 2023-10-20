package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrSellerStorerangeSync 同步商户中心服务范围
// tmall.nr.seller.storerange.sync
//
// 同步商户中心服务范围
func TmallNrSellerStorerangeSync(clt *core.SDKClient, req *tmallnr.TmallNrSellerStorerangeSyncAPIRequest, resp *tmallnr.TmallNrSellerStorerangeSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
