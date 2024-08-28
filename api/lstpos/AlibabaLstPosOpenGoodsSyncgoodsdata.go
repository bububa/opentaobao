package lstpos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenGoodsSyncgoodsdata 门店商品批量同步接口(最多10条商品信息)
// alibaba.lst.pos.open.goods.syncgoodsdata
//
// 门店商品批量同步接口(最多10条商品信息)
func AlibabaLstPosOpenGoodsSyncgoodsdata(ctx context.Context, clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest, resp *lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
