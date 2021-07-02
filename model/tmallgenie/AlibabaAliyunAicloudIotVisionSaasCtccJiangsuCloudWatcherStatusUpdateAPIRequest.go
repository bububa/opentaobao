package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 API请求
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
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

// NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest 初始化AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest对象
func NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest {
	return &AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Ctei Setter
// 设备唯一标识符
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// Get Ctei Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) GetCtei() string {
	return r._ctei
}

// Set is DevType Setter
// 设备对应的产品类型
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) SetDevType(_devType string) error {
	r._devType = _devType
	r.Set("dev_type", _devType)
	return nil
}

// Get DevType Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) GetDevType() string {
	return r._devType
}

// Set is SeqId Setter
// 一次请求的唯一标识符
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) SetSeqId(_seqId string) error {
	r._seqId = _seqId
	r.Set("seq_id", _seqId)
	return nil
}

// Get SeqId Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) GetSeqId() string {
	return r._seqId
}

// Set is UserAccount Setter
// 设备所属用户的账号信息
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) SetUserAccount(_userAccount string) error {
	r._userAccount = _userAccount
	r.Set("user_account", _userAccount)
	return nil
}

// Get UserAccount Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIRequest) GetUserAccount() string {
	return r._userAccount
}
