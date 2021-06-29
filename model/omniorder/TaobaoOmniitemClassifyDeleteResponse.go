package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一个分类 APIResponse
taobao.omniitem.classify.delete

删除一个分类
*/
type TaobaoOmniitemClassifyDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyDeleteResponse
}

type TaobaoOmniitemClassifyDeleteResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemClassifyDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
