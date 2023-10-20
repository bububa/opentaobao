package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangconsignorderbatchquery 发货单批量查询
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
func Alibabadchainaoxiangconsignorderbatchquery(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangconsignorderbatchqueryAPIRequest, session string) (*ascp.AlibabadchainaoxiangconsignorderbatchqueryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangconsignorderbatchqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
