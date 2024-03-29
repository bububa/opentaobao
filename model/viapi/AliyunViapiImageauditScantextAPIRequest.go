package viapi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImageauditScantextAPIRequest 文字内容审核 API请求
// aliyun.viapi.imageaudit.scantext
//
// 结合行为、内容，采用多维度、多模型、多检测手段，识别文本中的垃圾内容，规避色情、广告、灌水、渉政、辱骂等内容风险。
// 注意：如果返回结果里面的results为空，也代表指定类型检测通过。
type AliyunViapiImageauditScantextAPIRequest struct {
	model.Params
	// 指定检测对象，JSON数组中的每个元素是一个文件检测任务结构体（Task表）。最多支持10个元素，即对10张文本进行检测。每个元素的具体结构描述见Task。
	_tasks []Task
	// 指定文本检测的应用场景，可选值包括：  spam：含垃圾信息 politics： 涉政 abuse：辱骂 porn：智能鉴黄 terrorism：暴恐识别 flood：灌水 contraband：违禁 ad：文案违规识别 说明 支持多场景（Labels）一起检测，对一张文本同时进行鉴黄和暴恐识别，计费时也将按照两个场景计费。
	_labels []Label
}

// NewAliyunViapiImageauditScantextRequest 初始化AliyunViapiImageauditScantextAPIRequest对象
func NewAliyunViapiImageauditScantextRequest() *AliyunViapiImageauditScantextAPIRequest {
	return &AliyunViapiImageauditScantextAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunViapiImageauditScantextAPIRequest) Reset() {
	r._tasks = r._tasks[:0]
	r._labels = r._labels[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImageauditScantextAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageaudit.scantext"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImageauditScantextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiImageauditScantextAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTasks is Tasks Setter
// 指定检测对象，JSON数组中的每个元素是一个文件检测任务结构体（Task表）。最多支持10个元素，即对10张文本进行检测。每个元素的具体结构描述见Task。
func (r *AliyunViapiImageauditScantextAPIRequest) SetTasks(_tasks []Task) error {
	r._tasks = _tasks
	r.Set("tasks", _tasks)
	return nil
}

// GetTasks Tasks Getter
func (r AliyunViapiImageauditScantextAPIRequest) GetTasks() []Task {
	return r._tasks
}

// SetLabels is Labels Setter
// 指定文本检测的应用场景，可选值包括：  spam：含垃圾信息 politics： 涉政 abuse：辱骂 porn：智能鉴黄 terrorism：暴恐识别 flood：灌水 contraband：违禁 ad：文案违规识别 说明 支持多场景（Labels）一起检测，对一张文本同时进行鉴黄和暴恐识别，计费时也将按照两个场景计费。
func (r *AliyunViapiImageauditScantextAPIRequest) SetLabels(_labels []Label) error {
	r._labels = _labels
	r.Set("labels", _labels)
	return nil
}

// GetLabels Labels Getter
func (r AliyunViapiImageauditScantextAPIRequest) GetLabels() []Label {
	return r._labels
}

var poolAliyunViapiImageauditScantextAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunViapiImageauditScantextRequest()
	},
}

// GetAliyunViapiImageauditScantextRequest 从 sync.Pool 获取 AliyunViapiImageauditScantextAPIRequest
func GetAliyunViapiImageauditScantextAPIRequest() *AliyunViapiImageauditScantextAPIRequest {
	return poolAliyunViapiImageauditScantextAPIRequest.Get().(*AliyunViapiImageauditScantextAPIRequest)
}

// ReleaseAliyunViapiImageauditScantextAPIRequest 将 AliyunViapiImageauditScantextAPIRequest 放入 sync.Pool
func ReleaseAliyunViapiImageauditScantextAPIRequest(v *AliyunViapiImageauditScantextAPIRequest) {
	v.Reset()
	poolAliyunViapiImageauditScantextAPIRequest.Put(v)
}
