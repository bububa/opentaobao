package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除推广单元接口 APIResponse
taobao.simba.salestar.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaSalestarAdgroupDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarAdgroupDeleteResponse `json:"taobao_simba_salestar_adgroup_delete_response,omitempty"`
}

type TaobaoSimbaSalestarAdgroupDeleteResponse struct {

    // 删除失败时透出的原因，仅当success为false是有效
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 表示接口调用成功，当推广单元不存在等情况下也会返回true
    IsSuccess   bool `json:"is_success,omitempty"`

}
