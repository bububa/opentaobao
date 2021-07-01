package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

/* AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdate
天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 */
func AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdate(clt *core.SDKClient, req *tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest, session string) (*tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse, error) {
	var resp tmallgenie.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
