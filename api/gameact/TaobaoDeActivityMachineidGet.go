package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// Taobaodeactivitymachineidget 获取设备号
// taobao.de.activity.machineid.get
//
// 获取机器设备id
func Taobaodeactivitymachineidget(clt *core.SDKClient, req *gameact.TaobaodeactivitymachineidgetAPIRequest, session string) (*gameact.TaobaodeactivitymachineidgetAPIResponse, error) {
	var resp gameact.TaobaodeactivitymachineidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
