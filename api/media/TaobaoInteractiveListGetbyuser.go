package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaointeractivelistgetbyuser 用户获取视频互动列表
// taobao.interactive.list.getbyuser
//
// 根据用户来获取用户编辑的互动列表
func Taobaointeractivelistgetbyuser(clt *core.SDKClient, req *media.TaobaointeractivelistgetbyuserAPIRequest, session string) (*media.TaobaointeractivelistgetbyuserAPIResponse, error) {
	var resp media.TaobaointeractivelistgetbyuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
