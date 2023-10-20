package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoTmallgenieHotelplayerpause 天猫精灵酒店播放暂停
// taobao.tmallgenie.hotelplayerpause
//
// 酒店推送指令给天猫精灵停止播放音乐
func TaobaoTmallgenieHotelplayerpause(clt *core.SDKClient, req *tmallgenie.TaobaoTmallgenieHotelplayerpauseAPIRequest, resp *tmallgenie.TaobaoTmallgenieHotelplayerpauseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
