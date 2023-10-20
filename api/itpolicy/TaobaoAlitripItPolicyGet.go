package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// Taobaoalitripitpolicyget 【国际机票销售规则】单条查询
// taobao.alitrip.it.policy.get
//
// 通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号）
func Taobaoalitripitpolicyget(clt *core.SDKClient, req *itpolicy.TaobaoalitripitpolicygetAPIRequest, session string) (*itpolicy.TaobaoalitripitpolicygetAPIResponse, error) {
	var resp itpolicy.TaobaoalitripitpolicygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
