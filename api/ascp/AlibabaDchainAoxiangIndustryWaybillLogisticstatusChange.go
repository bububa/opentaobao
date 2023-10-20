package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangindustrywaybilllogisticstatuschange 物流节点回传
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.change
//
// 物流节点回传
func Alibabadchainaoxiangindustrywaybilllogisticstatuschange(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIRequest, session string) (*ascp.AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangindustrywaybilllogisticstatuschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
