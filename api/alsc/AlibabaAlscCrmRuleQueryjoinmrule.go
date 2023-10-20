package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequeryjoinmrule 查询品牌下的成为会员规则
// alibaba.alsc.crm.rule.queryjoinmrule
//
// 查询品牌下的成为会员规则
func Alibabaalsccrmrulequeryjoinmrule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequeryjoinmruleAPIRequest, session string) (*alsc.AlibabaalsccrmrulequeryjoinmruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequeryjoinmruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
