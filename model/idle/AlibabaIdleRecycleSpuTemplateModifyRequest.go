package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼接收回收商spu模板挂载信息 API请求
alibaba.idle.recycle.spu.template.modify

闲鱼接收回收商spu模板挂载信息
*/
type AlibabaIdleRecycleSpuTemplateModifyRequest struct {
    model.Params
    // 回收商挂载模版信息
    recycleSpuTemplate   *RecycleSpuTemplate
}

// 初始化AlibabaIdleRecycleSpuTemplateModifyRequest对象
func NewAlibabaIdleRecycleSpuTemplateModifyRequest() *AlibabaIdleRecycleSpuTemplateModifyRequest{
    return &AlibabaIdleRecycleSpuTemplateModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleSpuTemplateModifyRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.spu.template.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleSpuTemplateModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RecycleSpuTemplate Setter
// 回收商挂载模版信息
func (r *AlibabaIdleRecycleSpuTemplateModifyRequest) SetRecycleSpuTemplate(recycleSpuTemplate *RecycleSpuTemplate) error {
    r.recycleSpuTemplate = recycleSpuTemplate
    r.Set("recycle_spu_template", recycleSpuTemplate)
    return nil
}

// RecycleSpuTemplate Getter
func (r AlibabaIdleRecycleSpuTemplateModifyRequest) GetRecycleSpuTemplate() *RecycleSpuTemplate {
    return r.recycleSpuTemplate
}
