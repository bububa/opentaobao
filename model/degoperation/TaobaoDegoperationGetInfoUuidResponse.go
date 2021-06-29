package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid用户抽奖次数限制 APIResponse
taobao.degoperation.get.info.uuid

根据uuid用户抽奖次数限制
*/
type TaobaoDegoperationGetInfoUuidAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationGetInfoUuidResponse
}

type TaobaoDegoperationGetInfoUuidResponse struct {
    XMLName xml.Name `xml:"degoperation_get_info_uuid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
