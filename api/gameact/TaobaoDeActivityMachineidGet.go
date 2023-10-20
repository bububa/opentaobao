package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// TaobaoDeActivityMachineidGet 获取设备号
// taobao.de.activity.machineid.get
//
// 获取机器设备id
func TaobaoDeActivityMachineidGet(clt *core.SDKClient, req *gameact.TaobaoDeActivityMachineidGetAPIRequest, resp *gameact.TaobaoDeActivityMachineidGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
