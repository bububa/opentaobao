package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改创意与 APIResponse
taobao.simba.creative.update

更新一个创意的信息，可以设置创意标题、创意图片
*/
type TaobaoSimbaCreativeUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCreativeUpdateResponse `json:"taobao_simba_creative_update_response,omitempty"`
}

type TaobaoSimbaCreativeUpdateResponse struct {

    // 创意修改记录对象
    Creativerecord   *CreativeRecord `json:"creativerecord,omitempty"`

}
