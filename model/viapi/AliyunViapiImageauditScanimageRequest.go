package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绿网-内容安全 APIRequest
aliyun.viapi.imageaudit.scanimage

绿网-内容安全技术是基于阿里云视觉分析技术和深度识别技术，并经过在阿里经济体内和云上客户的多领域、多场景的广泛应用和不断优化，可提供风险和治理领域的图像识别、定位、检索等全面服务能力，不仅可以降低色情、涉恐、涉政、广告、垃圾信息等违规风险，而且能大幅度降低人工审核成本。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImageauditScanimageRequest struct {
    model.Params

    // 系统自动生成
    tasks   []Task 

    // 场景列表
    scenes   []string 

}

func NewAliyunViapiImageauditScanimageRequest() *AliyunViapiImageauditScanimageRequest{
    return &AliyunViapiImageauditScanimageRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunViapiImageauditScanimageRequest) GetApiMethodName() string {
    return "aliyun.viapi.imageaudit.scanimage"
}

func (r AliyunViapiImageauditScanimageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunViapiImageauditScanimageRequest) SetTasks(tasks []Task) error {
    r.tasks = tasks
    r.Set("tasks", tasks)
    return nil
}

func (r AliyunViapiImageauditScanimageRequest) GetTasks() []Task {
    return r.tasks
}

func (r *AliyunViapiImageauditScanimageRequest) SetScenes(scenes []string) error {
    r.scenes = scenes
    r.Set("scenes", scenes)
    return nil
}

func (r AliyunViapiImageauditScanimageRequest) GetScenes() []string {
    return r.scenes
}

