package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除adgroup的移动溢价 APIResponse
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_adgroup_mobilediscount_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回成功个数
    
    Result   int64 `json:"result,omitempty" xml:"