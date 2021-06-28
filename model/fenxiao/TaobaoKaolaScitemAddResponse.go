package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
考拉货品新增接口 APIResponse
taobao.kaola.scitem.add

考拉货品新增接口
*/
type TaobaoKaolaScitemAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"kaola_scitem_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异常信息
    
    ErrorMessages   []string `json:"error_messages,omitempty" xml:"