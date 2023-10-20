package choujiang

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/choujiang"
)

// Taobaodeactivitysecuritytokenapply 安全token获取
// taobao.de.activity.securitytoken.apply
//
// 新增接口，这个接口是用于在手机端进行抽奖时候的验证使用
func Taobaodeactivitysecuritytokenapply(clt *core.SDKClient, req *choujiang.TaobaodeactivitysecuritytokenapplyAPIRequest, session string) (*choujiang.TaobaodeactivitysecuritytokenapplyAPIResponse, error) {
	var resp choujiang.TaobaodeactivitysecuritytokenapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
