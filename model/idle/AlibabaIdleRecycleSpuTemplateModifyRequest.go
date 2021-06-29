package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼接收回收商spu模板挂载信息 APIRequest
alibaba.idle.recycle.spu.template.modify

闲鱼接收回收商spu模板挂载信息
*/
type AlibabaIdleRecycleSpuTemplateModifyRequest struct {
    model.Params

    // 回收商挂载模版信息
    recycleSpuTemplate   *RecycleSpuTemplate 

}

func NewAlibabaIdleRecycleSpuTemplateModifyRequest() *AlibabaIdleRecycleSpuTemplateModifyRequest{
    return &AlibabaIdleRecycleSpuTemplateModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRecycleSpuTemplateModifyRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.spu.template.modify"
}

func (r AlibabaIdleRecycleSpuTemplateModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRecycleSpuTemplateModifyRequest) SetRecycleSpuTemplate(recycleSpuTemplate *RecycleSpuTemplate) error {
    r.recycleSpuTemplate = recycleSpuTemplate
    r.Set("recycle_spu_template", recycleSpuTemplate)
    return nil
}

func (r AlibabaIdleRecycleSpuTemplateModifyRequest) GetRecycleSpuTemplate() *RecycleSpuTemplate {
    return r.recycleSpuTemplate
}

