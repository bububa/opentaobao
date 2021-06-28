package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）新建创意 APIResponse
taobao.simba.salestar.creative.add

创建一个创意
*/
type TaobaoSimbaSalestarCreativeAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_salestar_creative_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 新增加的创意对象
    
    Creative   *Creative `json:"creative,omitempty" xml:"