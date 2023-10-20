package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaotradedrugconfirmorder 阿里健康020接单
// taobao.trade.drug.confirmorder
//
// 阿里健康020接单
func Taobaotradedrugconfirmorder(clt *core.SDKClient, req *alihealth2.TaobaotradedrugconfirmorderAPIRequest, session string) (*alihealth2.TaobaotradedrugconfirmorderAPIResponse, error) {
	var resp alihealth2.TaobaotradedrugconfirmorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
