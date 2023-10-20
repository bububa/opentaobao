package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoInteractiveListGetbyuser 用户获取视频互动列表
// taobao.interactive.list.getbyuser
//
// 根据用户来获取用户编辑的互动列表
func TaobaoInteractiveListGetbyuser(clt *core.SDKClient, req *media.TaobaoInteractiveListGetbyuserAPIRequest, resp *media.TaobaoInteractiveListGetbyuserAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
