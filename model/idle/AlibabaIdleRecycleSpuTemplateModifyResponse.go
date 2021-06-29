package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼接收回收商spu模板挂载信息 API返回值 
alibaba.idle.recycle.spu.template.modify

闲鱼接收回收商spu模板挂载信息
*/
type AlibabaIdleRecycleSpuTemplateModifyAPIResponse struct {
    model.CommonResponse
    AlibabaIdleRecycleSpuTemplateModifyResponse
}

// 闲鱼接收回收商spu模板挂载信息 成功返回结果
type AlibabaIdleRecycleSpuTemplateModifyResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_recycle_spu_template_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *RecycleResult `json:"result,omitempty" xml:"result,omitempty"`
}
