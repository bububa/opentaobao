package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpPurchasePriceCreate ascp采购价写入接口
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
func AlibabaAscpPurchasePriceCreate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpPurchasePriceCreateAPIRequest, session string) (*ascpchannel.AlibabaAscpPurchasePriceCreateAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpPurchasePriceCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
