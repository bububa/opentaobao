package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdate 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
func AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdate(clt *core.SDKClient, req *tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest, resp *tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
