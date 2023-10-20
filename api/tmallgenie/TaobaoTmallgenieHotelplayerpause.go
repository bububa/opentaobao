package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaotmallgeniehotelplayerpause 天猫精灵酒店播放暂停
// taobao.tmallgenie.hotelplayerpause
//
// 酒店推送指令给天猫精灵停止播放音乐
func Taobaotmallgeniehotelplayerpause(clt *core.SDKClient, req *tmallgenie.TaobaotmallgeniehotelplayerpauseAPIRequest, session string) (*tmallgenie.TaobaotmallgeniehotelplayerpauseAPIResponse, error) {
	var resp tmallgenie.TaobaotmallgeniehotelplayerpauseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
