package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广组的信息 APIResponse
taobao.simba.adgroup.update

更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
*/
type TaobaoSimbaAdgroupUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_adgroup_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 被修改的推广组
    
    Adgroup   *ADGroup `json:"adgroup,omitempty" xml:"