package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpindustrylaunchextracharge 阿里巴巴.行业.增加费用.服务商发起
// alibaba.ascp.industry.launch.extra.charge
//
// 阿里巴巴.行业.增加费用.服务商发起
func Alibabaascpindustrylaunchextracharge(clt *core.SDKClient, req *ascpchannel.AlibabaascpindustrylaunchextrachargeAPIRequest, session string) (*ascpchannel.AlibabaascpindustrylaunchextrachargeAPIResponse, error) {
	var resp ascpchannel.AlibabaascpindustrylaunchextrachargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
