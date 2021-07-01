package choujiang

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/choujiang"
)

/* TaobaoDeActivitySecuritytokenApply
安全token获取
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用 */
func TaobaoDeActivitySecuritytokenApply(clt *core.SDKClient, req *choujiang.TaobaoDeActivitySecuritytokenApplyAPIRequest, session string) (*choujiang.TaobaoDeActivitySecuritytokenApplyAPIResponse, error) {
	var resp choujiang.TaobaoDeActivitySecuritytokenApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
