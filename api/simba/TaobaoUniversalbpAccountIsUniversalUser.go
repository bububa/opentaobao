package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpaccountisuniversaluser 判断用户是否迁移新bp
// taobao.universalbp.account.is.universal.user
//
// 获取客户是否应使用新接口。对于迁移了新bp的客户，使用新接口，没有迁移的，使用老bp接口。不可错乱使用。
func Taobaouniversalbpaccountisuniversaluser(clt *core.SDKClient, req *simba.TaobaouniversalbpaccountisuniversaluserAPIRequest, session string) (*simba.TaobaouniversalbpaccountisuniversaluserAPIResponse, error) {
	var resp simba.TaobaouniversalbpaccountisuniversaluserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
