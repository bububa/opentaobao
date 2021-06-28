package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户群列表 APIResponse
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表
*/
type TaobaoOpenimTribeGetalltribesAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_tribe_getalltribes_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 群列表信息
    
    TribeInfoList   []TribeInfo `json:"tribe_info_list,omitempty" xml:"