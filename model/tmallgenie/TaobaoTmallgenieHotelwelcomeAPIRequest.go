package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmallgenieHotelwelcomeAPIRequest
酒店欢迎词推送 API请求
taobao.tmallgenie.hotelwelcome

推送欢迎词，让天猫精灵播放 */
type TaobaoTmallgenieHotelwelcomeAPIRequest struct {
	model.Params
	// 房间号
	_roomNo string
	// 酒店ID
	_hotelId int64
	// 模板ID
	_templateId string
	// 模板变量
	_templateVariable string
}

// New
