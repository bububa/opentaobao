package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuanctguserrelation 用户
// alibaba.baichuan.ctg.user.relation
//
// 提供给优酷查询道长和淘宝账户的绑定关系
func Alibababaichuanctguserrelation(clt *core.SDKClient, req *baichuan.AlibababaichuanctguserrelationAPIRequest, session string) (*baichuan.AlibababaichuanctguserrelationAPIResponse, error) {
	var resp baichuan.AlibababaichuanctguserrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
