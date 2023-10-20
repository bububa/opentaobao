package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribelogsget openim 群聊天记录导出接口
// taobao.openim.tribelogs.get
//
// 获取openim账号的群聊天记录
func Taobaoopenimtribelogsget(clt *core.SDKClient, req *openim.TaobaoopenimtribelogsgetAPIRequest, session string) (*openim.TaobaoopenimtribelogsgetAPIResponse, error) {
	var resp openim.TaobaoopenimtribelogsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
