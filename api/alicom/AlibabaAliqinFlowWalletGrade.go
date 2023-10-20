package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowwalletgrade 获取流量档位
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
func Alibabaaliqinflowwalletgrade(clt *core.SDKClient, req *alicom.AlibabaaliqinflowwalletgradeAPIRequest, session string) (*alicom.AlibabaaliqinflowwalletgradeAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowwalletgradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
