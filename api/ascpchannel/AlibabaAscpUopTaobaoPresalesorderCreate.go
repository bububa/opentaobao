package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopTaobaoPresalesorderCreate 预售商家仓接单
// alibaba.ascp.uop.taobao.presalesorder.create
//
// 预售商家仓接单
func AlibabaAscpUopTaobaoPresalesorderCreate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest, session string) (*ascpchannel.AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpUopTaobaoPresalesorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
