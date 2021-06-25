package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户群列表 APIResponse
taobao.openim.tribe.getalltribes

OPENIM群服务获取用户群列表
*/
type TaobaoOpenimTribeGetalltribesAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimTribeGetalltribesResponse `json:"taobao_openim_tribe_getalltribes_response,omitempty"`
}

type TaobaoOpenimTribeGetalltribesResponse struct {

    // 群列表信息
    TribeInfoList   []TribeInfo `json:"tribe_info_list,omitempty"`

}
