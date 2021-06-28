package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取搜索溢价人群 APIResponse
taobao.simba.serchcrowd.get

根据推广单元id获取搜索溢价人群
*/
type TaobaoSimbaSerchcrowdGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_serchcrowd_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Adgrouptargetingtags   []TaobaoSimbaSerchcrowdGetResult `json:"adgrouptargetingtags,omitempty" xml:"