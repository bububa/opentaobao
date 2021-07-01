package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTaotvCarouselChannelAllAPIRequest
获取所有频道列表 API请求
taobao.taotv.carousel.channel.all

获取所有频道列表，按照序号升序 */
type TaobaoTaotvCarouselChannelAllAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// New
