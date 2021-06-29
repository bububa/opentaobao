package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_商品发布／更新 APIResponse
alibaba.alihealth.examination.goods.publish

体检机构对接_商品发布／更新
*/
type AlibabaAlihealthExaminationGoodsPublishAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationGoodsPublishResponse
}

type AlibabaAlihealthExaminationGoodsPublishResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_goods_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
