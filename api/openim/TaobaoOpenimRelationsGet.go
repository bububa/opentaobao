package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimRelationsGet 获取openim账号的聊天关系
// taobao.openim.relations.get
//
// 获取openim账号的聊天关系
func TaobaoOpenimRelationsGet(clt *core.SDKClient, req *openim.TaobaoOpenimRelationsGetAPIRequest, resp *openim.TaobaoOpenimRelationsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
