package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmallgenieHotelplayerpauseAPIRequest
天猫精灵酒店播放暂停 API请求
taobao.tmallgenie.hotelplayerpause

酒店推送指令给天猫精灵停止播放音乐 */
type TaobaoTmallgenieHotelplayerpauseAPIRequest struct {
	model.Params
	// 房间号
	_roomNo string
	// 酒店ID
	_hotelId int64
}

// New
