package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量修改单元智能出价 APIResponse
taobao.subway.cia.update

批量修改直通车推广单元的智能出价配置
*/
type TaobaoSubwayCiaUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSubwayCiaUpdateResponse
}

type TaobaoSubwayCiaUpdateResponse struct {
    XMLName xml.Name `xml:"subway_cia_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 推广组Id列表
    
    AdgroupList   []int64 `json:"adgroup_list,omitempty" xml:"adgroup_list>int64,omitempty"`
    
    
}
