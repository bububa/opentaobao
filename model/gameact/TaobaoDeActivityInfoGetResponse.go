package gameact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取活动信息 APIResponse
taobao.de.activity.info.get

根据appKey和活动id获取活动
*/
type TaobaoDeActivityInfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeActivityInfoGetResponse `json:"taobao_de_activity_info_get_response,omitempty"`
}

type TaobaoDeActivityInfoGetResponse struct {

    // 返回结构
    Activities   []ActivityVO `json:"activities,omitempty"`

}
