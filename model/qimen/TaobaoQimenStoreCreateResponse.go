package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店新增接口 APIResponse
taobao.qimen.store.create

isv调用接口来讲线下门店同步到线上
*/
type TaobaoQimenStoreCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStoreCreateResponse
}

type TaobaoQimenStoreCreateResponse struct {
    XMLName xml.Name `xml:"qimen_store_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 响应标示
    
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`

    
    // 响应code
    
    QimenCode   string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`

    
    // 返回的门店id
    
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`

    
}
