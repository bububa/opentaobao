package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询单元智能出价信息 APIResponse
taobao.subway.cia.get

查询单元智能出价信息
*/
type TaobaoSubwayCiaGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSubwayCiaGetResponse `json:"taobao_subway_cia_get_response,omitempty"`
}

type TaobaoSubwayCiaGetResponse struct {

    // 单元智能出价信息
    Result   *CiaConfig `json:"result,omitempty"`

}
