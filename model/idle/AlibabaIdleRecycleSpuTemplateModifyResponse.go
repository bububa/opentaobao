package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼接收回收商spu模板挂载信息 APIResponse
alibaba.idle.recycle.spu.template.modify

闲鱼接收回收商spu模板挂载信息
*/
type AlibabaIdleRecycleSpuTemplateModifyAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleSpuTemplateModifyResponse
}

type AlibabaIdleRecycleSpuTemplateModifyResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_spu_template_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RecycleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
