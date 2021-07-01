package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

/* TaobaoTmallgenieHotelwelcome
酒店欢迎词推送
taobao.tmallgenie.hotelwelcome

推送欢迎词，让天猫精灵播放 */
func TaobaoTmallgenieHotelwelcome(clt *core.SDKClient, req *tmallgenie.TaobaoTmallgenieHotelwelcomeAPIRequest, session string) (*tmallgenie.TaobaoTmallgenieHotelwelcomeAPIResponse, error) {
	var resp tmallgenie.TaobaoTmallgenieHotelwelcomeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
