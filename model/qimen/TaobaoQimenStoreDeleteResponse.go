package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店删除接口 APIResponse
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店
*/
type TaobaoQimenStoreDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_store_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应信息
    
    Message   string `json:"message,omitempty" xml:"