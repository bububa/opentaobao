package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// Taobaoalitripitpolicydelete 【国际机票销售规则】单条删除
// taobao.alitrip.it.policy.delete
//
// 销售规则删除接口，可以根据taobaoId或outId删除，根据outId删除时，如果outId不唯一，返回失败
func Taobaoalitripitpolicydelete(clt *core.SDKClient, req *itpolicy.TaobaoalitripitpolicydeleteAPIRequest, session string) (*itpolicy.TaobaoalitripitpolicydeleteAPIResponse, error) {
	var resp itpolicy.TaobaoalitripitpolicydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
