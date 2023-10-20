package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulelevelquerylevelrule 查询会员等级规则
// alibaba.alsc.crm.rule.level.querylevelrule
//
// 查询会员等级规则
func Alibabaalsccrmrulelevelquerylevelrule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulelevelquerylevelruleAPIRequest, session string) (*alsc.AlibabaalsccrmrulelevelquerylevelruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulelevelquerylevelruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
