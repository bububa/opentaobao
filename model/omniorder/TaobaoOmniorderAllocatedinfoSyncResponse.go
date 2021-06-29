package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分单结果同步给星盘 APIResponse
taobao.omniorder.allocatedinfo.sync

ISV分单完成，将分单结果同步给星盘
*/
type TaobaoOmniorderAllocatedinfoSyncAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderAllocatedinfoSyncResponse
}

type TaobaoOmniorderAllocatedinfoSyncResponse struct {
    XMLName xml.Name `xml:"omniorder_allocatedinfo_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`

    
    // 错误内容
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
