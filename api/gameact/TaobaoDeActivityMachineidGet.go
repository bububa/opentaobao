package gameact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

// TaobaoDeActivityMachineidGet 获取设备号
// taobao.de.activity.machineid.get
//
// 获取机器设备id
func TaobaoDeActivityMachineidGet(ctx context.Context, clt *core.SDKClient, req *gameact.TaobaoDeActivityMachineidGetAPIRequest, resp *gameact.TaobaoDeActivityMachineidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
