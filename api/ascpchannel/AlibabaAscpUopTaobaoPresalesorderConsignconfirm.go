package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopTaobaoPresalesorderConsignconfirm 预售商家仓出库
// alibaba.ascp.uop.taobao.presalesorder.consignconfirm
//
// 预售商家仓出库
func AlibabaAscpUopTaobaoPresalesorderConsignconfirm(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest, resp *ascpchannel.AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
