package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】基础信息获取接口：景点数据查询 APIResponse
taobao.alitrip.travel.baseinfo.scenics.get

商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
*/
type TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelBaseinfoScenicsGetResponse `json:"taobao_alitrip_travel_baseinfo_scenics_get_response,omitempty"`
}

type TaobaoAlitripTravelBaseinfoScenicsGetResponse struct {

    // 返回详细景点信息，返回数据为json数组结构的字符串
    ScenicInfos   string `json:"scenic_infos,omitempty"`

}
