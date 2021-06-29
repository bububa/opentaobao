package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店类目获取接口 APIResponse
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目
*/
type TaobaoQimenStorecategoryGetAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStorecategoryGetResponse
}

type TaobaoQimenStorecategoryGetResponse struct {
    XMLName xml.Name `xml:"qimen_storecategory_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应标示
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`

    
    // 类目json字符串
    
    StoreCategory   string `json:"store_category,omitempty" xml:"store_category,omitempty"`

    
    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应标示
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`

    
    // 类目json字符串
    
    StoreCategory   string `json:"store_category,omitempty" xml:"store_category,omitempty"`

    
}
