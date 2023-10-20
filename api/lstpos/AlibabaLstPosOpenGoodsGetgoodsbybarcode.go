package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// Alibabalstposopengoodsgetgoodsbybarcode ISV条码库查询接口
// alibaba.lst.pos.open.goods.getgoodsbybarcode
//
// ISV条码库查询接口
func Alibabalstposopengoodsgetgoodsbybarcode(clt *core.SDKClient, req *lstpos.AlibabalstposopengoodsgetgoodsbybarcodeAPIRequest, session string) (*lstpos.AlibabalstposopengoodsgetgoodsbybarcodeAPIResponse, error) {
	var resp lstpos.AlibabalstposopengoodsgetgoodsbybarcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
