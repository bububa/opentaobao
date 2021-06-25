package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取词包列表 APIResponse
taobao.subway.wordpackage.get

获取流量智选、捡漏词包等词包列表
*/
type TaobaoSubwayWordpackageGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSubwayWordpackageGetResponse `json:"taobao_subway_wordpackage_get_response,omitempty"`
}

type TaobaoSubwayWordpackageGetResponse struct {

    // 词包结果列表
    ResultList   []SiriusItemWordPackageDto `json:"result_list,omitempty"`

}
