package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
以盘点方式调整库存：传入商品实际库存 APIRequest
alibaba.mos.goods.synchinventorybycounting

以盘点方式调整库存：传入商品实际库存
盘点单自动判断数量增减
*/
type AlibabaMosGoodsSynchinventorybycountingRequest struct {
    model.Params

    // 专柜信息
    paramCountingInfoDto   *CountingInfoDto 

    // 盘点库存项（最大列表长度：20）
    countingItemDto   []CountingItemDto 

}

func NewAlibabaMosGoodsSynchinventorybycountingRequest() *AlibabaMosGoodsSynchinventorybycountingRequest{
    return &AlibabaMosGoodsSynchinventorybycountingRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.synchinventorybycounting"
}

func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosGoodsSynchinventorybycountingRequest) SetParamCountingInfoDto(paramCountingInfoDto *CountingInfoDto) error {
    r.paramCountingInfoDto = paramCountingInfoDto
    r.Set("param_counting_info_dto", paramCountingInfoDto)
    return nil
}

func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetParamCountingInfoDto() *CountingInfoDto {
    return r.paramCountingInfoDto
}

func (r *AlibabaMosGoodsSynchinventorybycountingRequest) SetCountingItemDto(countingItemDto []CountingItemDto) error {
    r.countingItemDto = countingItemDto
    r.Set("counting_item_dto", countingItemDto)
    return nil
}

func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetCountingItemDto() []CountingItemDto {
    return r.countingItemDto
}

