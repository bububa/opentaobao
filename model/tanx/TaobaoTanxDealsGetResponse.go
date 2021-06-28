package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取交易列表 APIResponse
taobao.tanx.deals.get

批量获取交易信息
*/
type TaobaoTanxDealsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_deals_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询是否成功
    
    Sucess   bool `json:"sucess,omitempty" xml:"