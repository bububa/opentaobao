package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加创意 APIResponse
taobao.simba.creative.add

创建一个创意
*/
type TaobaoSimbaCreativeAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCreativeAddResponse `json:"taobao_simba_creative_add_response,omitempty"`
}

type TaobaoSimbaCreativeAddResponse struct {

    // 新增加的创意对象
    Creative   *Creative `json:"creative,omitempty"`

}
