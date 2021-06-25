package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(新)创建一个推广组 APIResponse
taobao.simba.salestar.adgroup.add

创建一个推广组
*/
type TaobaoSimbaSalestarAdgroupAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarAdgroupAddResponse `json:"taobao_simba_salestar_adgroup_add_response,omitempty"`
}

type TaobaoSimbaSalestarAdgroupAddResponse struct {

    // 新增加的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
