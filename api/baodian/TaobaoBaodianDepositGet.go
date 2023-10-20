package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// Taobaobaodiandepositget 宝点用户帐户查询（已迁移）
// taobao.baodian.deposit.get
//
// 查询用户宝点帐户信息及当前宝点价格
func Taobaobaodiandepositget(clt *core.SDKClient, req *baodian.TaobaobaodiandepositgetAPIRequest, session string) (*baodian.TaobaobaodiandepositgetAPIResponse, error) {
	var resp baodian.TaobaobaodiandepositgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
