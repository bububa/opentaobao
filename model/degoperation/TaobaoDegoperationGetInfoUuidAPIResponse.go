package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据uuid用户抽奖次数限制 API返回值 
taobao.degoperation.get.info.uuid

根据uuid用户抽奖次数限制
*/
type TaobaoDegoperationGetInfoUuidAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationGetInfoUuidAPIResponseModel
}

// 根据uuid用户抽奖次数限制 成功返回结果
type TaobaoDegoperationGetInfoUuidAPIResponseModel struct {
    XMLName xml.Name `xml:"degoperation_get_info_uuid_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
