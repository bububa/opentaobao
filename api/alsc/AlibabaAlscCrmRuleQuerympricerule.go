package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequerympricerule 查询品牌下的会员价规则
// alibaba.alsc.crm.rule.querympricerule
//
// 查询品牌下的会员价规则
func Alibabaalsccrmrulequerympricerule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequerympriceruleAPIRequest, session string) (*alsc.AlibabaalsccrmrulequerympriceruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequerympriceruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
