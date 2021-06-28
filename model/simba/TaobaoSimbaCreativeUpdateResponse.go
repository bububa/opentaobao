package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改创意与 APIResponse
taobao.simba.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaCreativeUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_creative_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创意修改记录对象
    
    Creativerecord   *CreativeRecord `json:"creativerecord,omitempty" xml:"