package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单元智能出价信息 APIResponse
taobao.subway.cia.get

查询单元智能出价信息
*/
type TaobaoSubwayCiaGetAPIResponse struct {
    model.CommonResponse
    TaobaoSubwayCiaGetResponse
}

type TaobaoSubwayCiaGetResponse struct {
    XMLName xml.Name `xml:"subway_cia_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 单元智能出价信息
    
    Result   *CiaConfig `json:"result,omitempty" xml:"result,omitempty"`

    
}
