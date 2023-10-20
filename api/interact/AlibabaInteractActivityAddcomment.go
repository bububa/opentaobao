package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractactivityaddcomment 微淘评论接口
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
func Alibabainteractactivityaddcomment(clt *core.SDKClient, req *interact.AlibabainteractactivityaddcommentAPIRequest, session string) (*interact.AlibabainteractactivityaddcommentAPIResponse, error) {
	var resp interact.AlibabainteractactivityaddcommentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
