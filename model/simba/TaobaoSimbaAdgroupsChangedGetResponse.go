package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改的推广组ID和修改时间 APIResponse
taobao.simba.adgroups.changed.get

分页获取修改的推广组ID和修改时间
*/
type TaobaoSimbaAdgroupsChangedGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_adgroups_changed_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广组分页对象
    
    Adgroups   *ADGroupPage `json:"adgroups,omitempty" xml:"