package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanTaopasswordConfig 淘口令配置数据
// alibaba.baichuan.taopassword.config
//
// 百川淘口令规则配置接口
func AlibabaBaichuanTaopasswordConfig(clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordConfigAPIRequest, session string) (*baichuan.AlibabaBaichuanTaopasswordConfigAPIResponse, error) {
	var resp baichuan.AlibabaBaichuanTaopasswordConfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
