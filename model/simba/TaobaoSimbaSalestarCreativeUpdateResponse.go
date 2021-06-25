package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销量明星更新创意相关接口 APIResponse
taobao.simba.salestar.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaSalestarCreativeUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaSalestarCreativeUpdateResponse `json:"taobao_simba_salestar_creative_update_response,omitempty"`
}

type TaobaoSimbaSalestarCreativeUpdateResponse struct {

    // 创意修改记录对象
    Creativerecord   *CreativeRecord `json:"creativerecord,omitempty"`

}
