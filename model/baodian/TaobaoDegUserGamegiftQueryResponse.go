package baodian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户数娱游戏礼包查询 API返回值 
taobao.deg.user.gamegift.query

查询用户数娱礼包列表
*/
type TaobaoDegUserGamegiftQueryAPIResponse struct {
    model.CommonResponse
    TaobaoDegUserGamegiftQueryResponse
}

// 用户数娱游戏礼包查询 成功返回结果
type TaobaoDegUserGamegiftQueryResponse struct {
    XMLName xml.Name `xml:"deg_user_gamegift_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 礼包信息
    Records   []GameGiftRecordDto `json:"records,omitempty" xml:"records>game_gift_record_dto,omitempty"`
}
