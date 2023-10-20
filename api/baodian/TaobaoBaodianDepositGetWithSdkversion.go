package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// Taobaobaodiandepositgetwithsdkversion 查询用户宝点信息（带sdk版本，已迁移）
// taobao.baodian.deposit.get.with.sdkversion
//
// 获取用户宝点信息（带sdk版本，已迁移）
func Taobaobaodiandepositgetwithsdkversion(clt *core.SDKClient, req *baodian.TaobaobaodiandepositgetwithsdkversionAPIRequest, session string) (*baodian.TaobaobaodiandepositgetwithsdkversionAPIResponse, error) {
	var resp baodian.TaobaobaodiandepositgetwithsdkversionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
