package gameact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取活动信息 API返回值 
taobao.de.activity.info.get

根据appKey和活动id获取活动
*/
type TaobaoDeActivityInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoDeActivityInfoGetAPIResponseModel
}

// 获取活动信息 成功返回结果
type TaobaoDeActivityInfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"de_activity_info_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结构
    Activities   []ActivityVO `json:"activities,omitempty" xml:"activities>activity_vo,omitempty"`
}
