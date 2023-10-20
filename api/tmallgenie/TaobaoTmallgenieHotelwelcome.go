package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaotmallgeniehotelwelcome 酒店欢迎词推送
// taobao.tmallgenie.hotelwelcome
//
// 推送欢迎词，让天猫精灵播放
func Taobaotmallgeniehotelwelcome(clt *core.SDKClient, req *tmallgenie.TaobaotmallgeniehotelwelcomeAPIRequest, session string) (*tmallgenie.TaobaotmallgeniehotelwelcomeAPIResponse, error) {
	var resp tmallgenie.TaobaotmallgeniehotelwelcomeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
