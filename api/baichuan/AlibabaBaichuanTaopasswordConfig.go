package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuantaopasswordconfig 淘口令配置数据
// alibaba.baichuan.taopassword.config
//
// 百川淘口令规则配置接口
func Alibababaichuantaopasswordconfig(clt *core.SDKClient, req *baichuan.AlibababaichuantaopasswordconfigAPIRequest, session string) (*baichuan.AlibababaichuantaopasswordconfigAPIResponse, error) {
	var resp baichuan.AlibababaichuantaopasswordconfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
