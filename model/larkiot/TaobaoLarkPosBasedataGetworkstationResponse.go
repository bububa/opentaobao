package larkiot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据影城id工作站和macId获取工作站 APIResponse
taobao.lark.pos.basedata.getworkstation

获取单独工作站
*/
type TaobaoLarkPosBasedataGetworkstationAPIResponse struct {
    model.CommonResponse
    TaobaoLarkPosBasedataGetworkstationResponse
}

type TaobaoLarkPosBasedataGetworkstationResponse struct {
    XMLName xml.Name `xml:"lark_pos_basedata_getworkstation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应结果
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
}
