package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest
天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 API请求
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 */
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest struct {
	model.Params
	// 设备唯一标识符
	_ctei string
	// 设备对应的产品类型
	_devType string
	// 一次请求的唯一标识符
	_seqId string
	// 设备所属用户的账号信息
	_userAccount string
}

// New
