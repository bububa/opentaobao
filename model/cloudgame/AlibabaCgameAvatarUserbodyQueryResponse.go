package cloudgame

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户Avatar body查询 APIResponse
alibaba.cgame.avatar.userbody.query

Avatar用户body数据查询
*/
type AlibabaCgameAvatarUserbodyQueryAPIResponse struct {
    model.CommonResponse
    AlibabaCgameAvatarUserbodyQueryResponse
}

type AlibabaCgameAvatarUserbodyQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_cgame_avatar_userbody_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaCgameAvatarUserbodyQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
