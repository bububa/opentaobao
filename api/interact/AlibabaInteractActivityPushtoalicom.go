package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractActivityPushtoalicom 小铺isv推广流量活动到流量聚乐部
// alibaba.interact.activity.pushtoalicom
//
// 涉及到流量包的小铺isv，将活动推送到流量聚乐部
func AlibabaInteractActivityPushtoalicom(clt *core.SDKClient, req *interact.AlibabaInteractActivityPushtoalicomAPIRequest, resp *interact.AlibabaInteractActivityPushtoalicomAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
