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
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_store_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"