package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopskilslist 获取硬件平台设备下挂载的技能列表
// taobao.ailab.aicloud.top.skils.list
//
// 提供给在硬件平台接入设备的技能列表
func Taobaoailabaicloudtopskilslist(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopskilslistAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopskilslistAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopskilslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
