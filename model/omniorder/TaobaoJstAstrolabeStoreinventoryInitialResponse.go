package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
后端商品库存初始化 APIResponse
taobao.jst.astrolabe.storeinventory.initial

初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。
*/
type TaobaoJstAstrolabeStoreinventoryInitialAPIResponse struct {
    model.CommonResponse
    TaobaoJstAstrolabeStoreinventoryInitialResponse
}

type TaobaoJstAstrolabeStoreinventoryInitialResponse struct {
    XMLName xml.Name `xml:"jst_astrolabe_storeinventory_initial_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应标签
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 响应标示
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // 错误信息列表
    
    ErrorDescriptions   []Error `json:"error_descriptions,omitempty" xml:"error_descriptions>error,omitempty"`
    
    
}
