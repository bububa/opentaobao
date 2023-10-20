package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuanasoactivate 设备安装活动激活
// alibaba.baichuan.aso.activate
//
// 设备安装活动激活
func Alibababaichuanasoactivate(clt *core.SDKClient, req *baichuan.AlibababaichuanasoactivateAPIRequest, session string) (*baichuan.AlibababaichuanasoactivateAPIResponse, error) {
	var resp baichuan.AlibababaichuanasoactivateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
