package yunosad

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosad"
)

// YunosAdAuditCreativeGet 获取单个创意审核状态
// yunos.ad.audit.creative.get
//
// 获取单个创意审核状态
func YunosAdAuditCreativeGet(clt *core.SDKClient, req *yunosad.YunosAdAuditCreativeGetAPIRequest, resp *yunosad.YunosAdAuditCreativeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
