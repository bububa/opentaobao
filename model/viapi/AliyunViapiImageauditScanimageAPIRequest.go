package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImageauditScanimageAPIRequest 绿网-内容安全 API请求
// aliyun.viapi.imageaudit.scanimage
//
// 绿网-内容安全技术是基于阿里云视觉分析技术和深度识别技术，并经过在阿里经济体内和云上客户的多领域、多场景的广泛应用和不断优化，可提供风险和治理领域的图像识别、定位、检索等全面服务能力，不仅可以降低色情、涉恐、涉政、广告、垃圾信息等违规风险，而且能大幅度降低人工审核成本。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImageauditScanimageAPIRequest struct {
	model.Params
	// 系统自动生成
	_tasks []Task
	// 场景列表
	_scenes []string
}

// NewAliyunViapiImageauditScanimageRequest 初始化AliyunViapiImageauditScanimageAPIRequest对象
func NewAliyunViapiImageauditScanimageRequest() *AliyunViapiImageauditScanimageAPIRequest {
	return &AliyunViapiImageauditScanimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImageauditScanimageAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageaudit.scanimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImageauditScanimageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTasks is Tasks Setter
// 系统自动生成
func (r *AliyunViapiImageauditScanimageAPIRequest) SetTasks(_tasks []Task) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// GetTasks Tasks Getter
func (r AliyunViapiImageauditScanimageAPIRequest) GetTasks() []Task {
	return r._tasks
}

// SetScenes is Scenes Setter
// 场景列表
func (r *AliyunViapiImageauditScanimageAPIRequest) SetScenes(_scenes []string) error {
	r._scenes = _scenes
	r.Set("scenes", _scenes)
	return nil
}

// GetScenes Scenes Getter
func (r AliyunViapiImageauditScanimageAPIRequest) GetScenes() []string {
	return r._scenes
}
