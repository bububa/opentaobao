package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractActivityAddcomment 微淘评论接口
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
func AlibabaInteractActivityAddcomment(clt *core.SDKClient, req *interact.AlibabaInteractActivityAddcommentAPIRequest, resp *interact.AlibabaInteractActivityAddcommentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
