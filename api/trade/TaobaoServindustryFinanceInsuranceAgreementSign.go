package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoservindustryfinanceinsuranceagreementsign 保司合同签约后回调接口
// taobao.servindustry.finance.insurance.agreement.sign
//
// 保司合同签约后回调接口
func Taobaoservindustryfinanceinsuranceagreementsign(clt *core.SDKClient, req *trade.TaobaoservindustryfinanceinsuranceagreementsignAPIRequest, session string) (*trade.TaobaoservindustryfinanceinsuranceagreementsignAPIResponse, error) {
	var resp trade.TaobaoservindustryfinanceinsuranceagreementsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
