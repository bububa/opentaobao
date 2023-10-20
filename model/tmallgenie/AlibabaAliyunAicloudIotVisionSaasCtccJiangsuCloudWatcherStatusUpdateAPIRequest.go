package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 API请求
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
type AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest struct {
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

// NewAlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateRequest 初始化AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest对象
func NewAlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateRequest() *AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest {
	return &AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCtei is Ctei Setter
// 设备唯一标识符
func (r *AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) SetCtei(_ctei string) error {
	r._ctei = _ctei
	r.Set("ctei", _ctei)
	return nil
}

// GetCtei Ctei Getter
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetCtei() string {
	return r._ctei
}

// SetDevType is DevType Setter
// 设备对应的产品类型
func (r *AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) SetDevType(_devType string) error {
	r._devType = _devType
	r.Set("dev_type", _devType)
	return nil
}

// GetDevType DevType Getter
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetDevType() string {
	return r._devType
}

// SetSeqId is SeqId Setter
// 一次请求的唯一标识符
func (r *AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) SetSeqId(_seqId string) error {
	r._seqId = _seqId
	r.Set("seq_id", _seqId)
	return nil
}

// GetSeqId SeqId Getter
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetSeqId() string {
	return r._seqId
}

// SetUserAccount is UserAccount Setter
// 设备所属用户的账号信息
func (r *AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) SetUserAccount(_userAccount string) error {
	r._userAccount = _userAccount
	r.Set("user_account", _userAccount)
	return nil
}

// GetUserAccount UserAccount Getter
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsucloudwatcherstatusupdateAPIRequest) GetUserAccount() string {
	return r._userAccount
}
