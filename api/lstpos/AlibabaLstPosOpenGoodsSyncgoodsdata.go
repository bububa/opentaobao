package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenGoodsSyncgoodsdata 门店商品批量同步接口(最多10条商品信息)
// alibaba.lst.pos.open.goods.syncgoodsdata
//
// 门店商品批量同步接口(最多10条商品信息)
func AlibabaLstPosOpenGoodsSyncgoodsdata(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest, resp *lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
