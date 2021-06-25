package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广组 APIResponse
taobao.simba.adgroupsbyadgroupids.get

批量得到推广组
*/
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupsbyadgroupidsGetResponse `json:"taobao_simba_adgroupsbyadgroupids_get_response,omitempty"`
}

type TaobaoSimbaAdgroupsbyadgroupidsGetResponse struct {

    // 返回的推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty"`

}
