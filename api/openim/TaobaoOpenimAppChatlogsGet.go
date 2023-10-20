package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimappchatlogsget openim应用聊天记录查询
// taobao.openim.app.chatlogs.get
//
// 查询openim应用的聊天记录
func Taobaoopenimappchatlogsget(clt *core.SDKClient, req *openim.TaobaoopenimappchatlogsgetAPIRequest, session string) (*openim.TaobaoopenimappchatlogsgetAPIResponse, error) {
	var resp openim.TaobaoopenimappchatlogsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
