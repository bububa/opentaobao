package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimchatlogsget openim聊天记录查询接口
// taobao.openim.chatlogs.get
//
// 查询openim账号聊天记录
func Taobaoopenimchatlogsget(clt *core.SDKClient, req *openim.TaobaoopenimchatlogsgetAPIRequest, session string) (*openim.TaobaoopenimchatlogsgetAPIResponse, error) {
	var resp openim.TaobaoopenimchatlogsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
