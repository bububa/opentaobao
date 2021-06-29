package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文字内容审核 API请求
aliyun.viapi.imageaudit.scantext

结合行为、内容，采用多维度、多模型、多检测手段，识别文本中的垃圾内容，规避色情、广告、灌水、渉政、辱骂等内容风险。
注意：如果返回结果里面的results为空，也代表指定类型检测通过。
*/
type AliyunViapiImageauditScantextRequest struct {
    model.Params
    // 指定检测对象，JSON数组中的每个元素是一个文件检测任务结构体（Task表）。最多支持10个元素，即对10张文本进行检测。每个元素的具体结构描述见Task。
    _tasks   []Task
    // 指定文本检测的应用场景，可选值包括：  spam：含垃圾信息 politics： 涉政 abuse：辱骂 porn：智能鉴黄 terrorism：暴恐识别 flood：灌水 contraband：违禁 ad：文案违规识别 说明 支持多场景（Labels）一起检测，对一张文本同时进行鉴黄和暴恐识别，计费时也将按照两个场景计费。
    _labels   []Label
}

// 初始化AliyunViapiImageauditScantextRequest对象
func NewAliyunViapiImageauditScantextRequest() *AliyunViapiImageauditScantextRequest{
    return &AliyunViapiImageauditScantextRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiImageauditScantextRequest) GetApiMethodName() string {
    return "aliyun.viapi.imageaudit.scantext"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiImageauditScantextRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tasks Setter
// 指定检测对象，JSON数组中的每个元素是一个文件检测任务结构体（Task表）。最多支持10个元素，即对10张文本进行检测。每个元素的具体结构描述见Task。
func (r *AliyunViapiImageauditScantextRequest) SetTasks(_tasks []Task) error {
    r._tasks = _tasks
    r.Set("tasks", _tasks)
    return nil
}

// Tasks Getter
func (r AliyunViapiImageauditScantextRequest) GetTasks() []Task {
    return r._tasks
}
// Labels Setter
// 指定文本检测的应用场景，可选值包括：  spam：含垃圾信息 politics： 涉政 abuse：辱骂 porn：智能鉴黄 terrorism：暴恐识别 flood：灌水 contraband：违禁 ad：文案违规识别 说明 支持多场景（Labels）一起检测，对一张文本同时进行鉴黄和暴恐识别，计费时也将按照两个场景计费。
func (r *AliyunViapiImageauditScantextRequest) SetLabels(_labels []Label) error {
    r._labels = _labels
    r.Set("labels", _labels)
    return nil
}

// Labels Getter
func (r AliyunViapiImageauditScantextRequest) GetLabels() []Label {
    return r._labels
}
