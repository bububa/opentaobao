package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopTaobaoWaybillQuerycp 查询电子面单开放的CP列表
// alibaba.ascp.uop.taobao.waybill.querycp
//
// 查询电子面单开放的CP列表
func AlibabaAscpUopTaobaoWaybillQuerycp(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest, session string) (*ascpchannel.AlibabaAscpUopTaobaoWaybillQuerycpAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpUopTaobaoWaybillQuerycpAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
