package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品级联属性信息获取 APIResponse
alibaba.item.publish.props.get

新商品发布，商品级联属性信息获取
*/
type AlibabaItemPublishPropsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_item_publish_props_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品发布规则信息，XML格式.
    
    Result   string `json:"result,omitempty" xml:"