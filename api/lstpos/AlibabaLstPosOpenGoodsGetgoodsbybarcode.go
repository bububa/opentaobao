package lstpos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenGoodsGetgoodsbybarcode ISV条码库查询接口
// alibaba.lst.pos.open.goods.getgoodsbybarcode
//
// ISV条码库查询接口
func AlibabaLstPosOpenGoodsGetgoodsbybarcode(ctx context.Context, clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest, resp *lstpos.AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
