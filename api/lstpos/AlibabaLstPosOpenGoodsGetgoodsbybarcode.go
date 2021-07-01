package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

/* AlibabaLstPosOpenGoodsGetgoodsbybarcode
ISV条码库查询接口
alibaba.lst.pos.open.goods.getgoodsbybarcode

ISV条码库查询接口 */
func AlibabaLstPosOpenGoodsGetgoodsbybarcode(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIRequest, session string) (*lstpos.AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIResponse, error) {
	var resp lstpos.AlibabaLstPosOpenGoodsGetgoodsbybarcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
