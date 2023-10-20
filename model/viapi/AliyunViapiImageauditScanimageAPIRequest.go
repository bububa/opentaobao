package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapiimageauditscanimageAPIRequest 绿网-内容安全 API请求
// aliyun.viapi.imageaudit.scanimage
//
// 绿网-内容安全技术是基于阿里云视觉分析技术和深度识别技术，并经过在阿里经济体内和云上客户的多领域、多场景的广泛应用和不断优化，可提供风险和治理领域的图像识别、定位、检索等全面服务能力，不仅可以降低色情、涉恐、涉政、广告、垃圾信息等违规风险，而且能大幅度降低人工审核成本。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapiimageauditscanimageAPIRequest struct {
	model.Params
	// 系统自动生成
	_tasks []Task
	// 场景列表
	_scenes []string
}

// NewAliyunviapiimageauditscanimageRequest 初始化AliyunviapiimageauditscanimageAPIRequest对象
func NewAliyunviapiimageauditscanimageRequest() *AliyunviapiimageauditscanimageAPIRequest {
	return &AliyunviapiimageauditscanimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapiimageauditscanimageAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageaudit.scanimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapiimageauditscanimageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapiimageauditscanimageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTasks is Tasks Setter
// 系统自动生成
func (r *AliyunviapiimageauditscanimageAPIRequest) SetTasks(_tasks []Task) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// GetTasks Tasks Getter
func (r AliyunviapiimageauditscanimageAPIRequest) GetTasks() []Task {
	return r._tasks
}

// SetScenes is Scenes Setter
// 场景列表
func (r *AliyunviapiimageauditscanimageAPIRequest) SetScenes(_scenes []string) error {
	r._scenes = _scenes
	r.Set("scenes", _scenes)
	return nil
}

// GetScenes Scenes Getter
func (r AliyunviapiimageauditscanimageAPIRequest) GetScenes() []string {
	return r._scenes
}
