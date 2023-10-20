package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Alibabaalihealthnrtrademedicalinsuranceget 阿里健康医保支付信息获取
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
func Alibabaalihealthnrtrademedicalinsuranceget(clt *core.SDKClient, req *drug.AlibabaalihealthnrtrademedicalinsurancegetAPIRequest, session string) (*drug.AlibabaalihealthnrtrademedicalinsurancegetAPIResponse, error) {
	var resp drug.AlibabaalihealthnrtrademedicalinsurancegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
