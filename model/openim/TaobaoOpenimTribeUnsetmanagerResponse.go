package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群取消管理员 API返回值 
taobao.openim.tribe.unsetmanager

OPENIM群取消管理员
*/
type TaobaoOpenimTribeUnsetmanagerAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeUnsetmanagerResponse
}

// OPENIM群取消管理员 成功返回结果
type TaobaoOpenimTribeUnsetmanagerResponse struct {
    XMLName xml.Name `xml:"openim_tribe_unsetmanager_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 群服务code
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
