package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）新建创意 APIResponse
taobao.simba.salestar.creative.add

创建一个创意
*/
type TaobaoSimbaSalestarCreativeAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarCreativeAddResponse `json:"taobao_simba_salestar_creative_add_response,omitempty"`
}

type TaobaoSimbaSalestarCreativeAddResponse struct {

    // 新增加的创意对象
    Creative   *Creative `json:"creative,omitempty"`

}
