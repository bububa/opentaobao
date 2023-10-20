package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// Alibabalstposopengoodssyncgoodsdata 门店商品批量同步接口(最多10条商品信息)
// alibaba.lst.pos.open.goods.syncgoodsdata
//
// 门店商品批量同步接口(最多10条商品信息)
func Alibabalstposopengoodssyncgoodsdata(clt *core.SDKClient, req *lstpos.AlibabalstposopengoodssyncgoodsdataAPIRequest, session string) (*lstpos.AlibabalstposopengoodssyncgoodsdataAPIResponse, error) {
	var resp lstpos.AlibabalstposopengoodssyncgoodsdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
