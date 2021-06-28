package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查看功能权限 APIResponse
taobao.simba.customers.sid.get

查询用户是否拥有某个功能权限
*/
type TaobaoSimbaCustomersSidGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_customers_sid_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 权限列表及是否有权限
    
    Result   *SidVo `json:"result,omitempty" xml:"