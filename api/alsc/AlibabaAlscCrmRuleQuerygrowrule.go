package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequerygrowrule 查询品牌下的会员成长规则
// alibaba.alsc.crm.rule.querygrowrule
//
// 查询品牌下的会员成长规则
func Alibabaalsccrmrulequerygrowrule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequerygrowruleAPIRequest, session string) (*alsc.AlibabaalsccrmrulequerygrowruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequerygrowruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
