package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量更新词包 APIResponse
taobao.subway.wordpackage.update

批量更新词包
*/
type TaobaoSubwayWordpackageUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSubwayWordpackageUpdateResponse `json:"taobao_subway_wordpackage_update_response,omitempty"`
}

type TaobaoSubwayWordpackageUpdateResponse struct {

    // 接口返回model
    Result   *TaobaoSubwayWordpackageUpdateResult `json:"result,omitempty"`

}
