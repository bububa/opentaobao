package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractSensorSocial 社交组件
// alibaba.interact.sensor.social
//
// 赞，评论 ，关注 新增接口
func AlibabaInteractSensorSocial(clt *core.SDKClient, req *mtopopen.AlibabaInteractSensorSocialAPIRequest, resp *mtopopen.AlibabaInteractSensorSocialAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
