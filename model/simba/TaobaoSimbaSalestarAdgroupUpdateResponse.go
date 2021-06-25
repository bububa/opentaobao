package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销量明星更新一个推广组的信息 APIResponse
taobao.simba.salestar.adgroup.update

更新一个推广组的信息，可以设置 是否上线
*/
type TaobaoSimbaSalestarAdgroupUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarAdgroupUpdateResponse `json:"taobao_simba_salestar_adgroup_update_response,omitempty"`
}

type TaobaoSimbaSalestarAdgroupUpdateResponse struct {

    // 被修改的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
