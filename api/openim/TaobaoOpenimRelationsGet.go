package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimrelationsget 获取openim账号的聊天关系
// taobao.openim.relations.get
//
// 获取openim账号的聊天关系
func Taobaoopenimrelationsget(clt *core.SDKClient, req *openim.TaobaoopenimrelationsgetAPIRequest, session string) (*openim.TaobaoopenimrelationsgetAPIResponse, error) {
	var resp openim.TaobaoopenimrelationsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
