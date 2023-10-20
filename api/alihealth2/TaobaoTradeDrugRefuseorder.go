package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Taobaotradedrugrefuseorder 阿里健康020拒单
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
func Taobaotradedrugrefuseorder(clt *core.SDKClient, req *alihealth2.TaobaotradedrugrefuseorderAPIRequest, session string) (*alihealth2.TaobaotradedrugrefuseorderAPIResponse, error) {
	var resp alihealth2.TaobaotradedrugrefuseorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
