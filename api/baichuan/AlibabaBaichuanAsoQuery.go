package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuanasoquery 查询app在设备上的安装信息
// alibaba.baichuan.aso.query
//
// 查询app在设备上的安装信息
func Alibababaichuanasoquery(clt *core.SDKClient, req *baichuan.AlibababaichuanasoqueryAPIRequest, session string) (*baichuan.AlibababaichuanasoqueryAPIResponse, error) {
	var resp baichuan.AlibababaichuanasoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
