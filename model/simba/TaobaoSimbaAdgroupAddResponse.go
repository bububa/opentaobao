package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广组 APIResponse
taobao.simba.adgroup.add

创建一个推广组
*/
type TaobaoSimbaAdgroupAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupAddResponse `json:"taobao_simba_adgroup_add_response,omitempty"`
}

type TaobaoSimbaAdgroupAddResponse struct {

    // 新增加的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
