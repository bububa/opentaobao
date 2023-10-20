package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangindustrywaybilllogisticstatusremarkadd 客服增加备注
// alibaba.dchain.aoxiang.industry.waybill.logisticstatus.remark.add
//
// 客服增加备注
func Alibabadchainaoxiangindustrywaybilllogisticstatusremarkadd(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIRequest, session string) (*ascp.AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangindustrywaybilllogisticstatusremarkaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
