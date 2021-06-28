package baodian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户数娱游戏礼包查询 APIResponse
taobao.deg.user.gamegift.query

查询用户数娱礼包列表
*/
type TaobaoDegUserGamegiftQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"deg_user_gamegift_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 礼包信息
    
    Records   []GameGiftRecordDto `json:"records,omitempty" xml:"