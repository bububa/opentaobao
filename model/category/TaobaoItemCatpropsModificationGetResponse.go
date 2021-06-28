package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品类目属性变更 APIResponse
taobao.item.catprops.modification.get

查询商品类目属性变更信息
*/
type TaobaoItemCatpropsModificationGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemCatpropsModificationGetResponse
}

type TaobaoItemCatpropsModificationGetResponse struct {
    XMLName xml.Name `xml:"item_catprops_modification_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Results   []PropsModificationResult `json:"results,omitempty" xml:"results>props_modification_result,omitempty"`
    
    
}
