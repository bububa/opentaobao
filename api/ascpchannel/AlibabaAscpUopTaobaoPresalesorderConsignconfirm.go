package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopTaobaoPresalesorderConsignconfirm 预售商家仓出库
// alibaba.ascp.uop.taobao.presalesorder.consignconfirm
//
// 预售商家仓出库
func AlibabaAscpUopTaobaoPresalesorderConsignconfirm(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest, resp *ascpchannel.AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
