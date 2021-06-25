package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获得创意 APIResponse
taobao.simba.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
type TaobaoSimbaCreativesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCreativesGetResponse `json:"taobao_simba_creatives_get_response,omitempty"`
}

type TaobaoSimbaCreativesGetResponse struct {

    // 创意对象列表
    Creatives   []Creative `json:"creatives,omitempty"`

}
