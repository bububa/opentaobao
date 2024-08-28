package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrSellerStorerangeRead 门店服务范围读取
// tmall.nr.seller.storerange.read
//
// 读取卖家所属门店的服务范围
func TmallNrSellerStorerangeRead(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrSellerStorerangeReadAPIRequest, resp *tmallnr.TmallNrSellerStorerangeReadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
