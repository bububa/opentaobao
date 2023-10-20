package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Cainiaobimtradeorderconsign 驱动保税交易订单发货
// cainiao.bim.tradeorder.consign
//
// 驱动保税交易订单发货
func Cainiaobimtradeorderconsign(clt *core.SDKClient, req *wms.CainiaobimtradeorderconsignAPIRequest, session string) (*wms.CainiaobimtradeorderconsignAPIResponse, error) {
	var resp wms.CainiaobimtradeorderconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
