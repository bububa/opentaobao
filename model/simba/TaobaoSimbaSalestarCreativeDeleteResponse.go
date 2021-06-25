package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除创意相关接口 APIResponse
taobao.simba.salestar.creative.delete

删除一个创意
*/
type TaobaoSimbaSalestarCreativeDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarCreativeDeleteResponse `json:"taobao_simba_salestar_creative_delete_response,omitempty"`
}

type TaobaoSimbaSalestarCreativeDeleteResponse struct {

    // 被删除的创意对象
    Creative   *Creative `json:"creative,omitempty"`

}
