package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrSellerStorerangeRead 门店服务范围读取
// tmall.nr.seller.storerange.read
//
// 读取卖家所属门店的服务范围
func TmallNrSellerStorerangeRead(clt *core.SDKClient, req *tmallnr.TmallNrSellerStorerangeReadAPIRequest, resp *tmallnr.TmallNrSellerStorerangeReadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
