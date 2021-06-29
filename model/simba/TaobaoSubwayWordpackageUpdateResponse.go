package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新词包 APIResponse
taobao.subway.wordpackage.update

批量更新词包
*/
type TaobaoSubwayWordpackageUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSubwayWordpackageUpdateResponse
}

type TaobaoSubwayWordpackageUpdateResponse struct {
    XMLName xml.Name `xml:"subway_wordpackage_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoSubwayWordpackageUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
