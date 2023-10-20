package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemSubscribeRelationQuery 查询单个订阅关系
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
func TaobaoBaichuanItemSubscribeRelationQuery(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeRelationQueryAPIRequest, resp *baichuan.TaobaoBaichuanItemSubscribeRelationQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
