package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpIndustryLogisticsSync 物流状态同步
// alibaba.ascp.industry.logistics.sync
//
// 履约物流状态同步
func AlibabaAscpIndustryLogisticsSync(clt *core.SDKClient, req *ascpchannel.AlibabaAscpIndustryLogisticsSyncAPIRequest, session string) (*ascpchannel.AlibabaAscpIndustryLogisticsSyncAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpIndustryLogisticsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
