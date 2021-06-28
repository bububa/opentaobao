package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设置ios证书 APIResponse
taobao.openim.ioscert.production.set

设置ios证书
*/
type TaobaoOpenimIoscertProductionSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_ioscert_production_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作成功
    
    Code   string `json:"code,omitempty" xml:"