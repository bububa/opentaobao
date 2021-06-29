package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
差旅申请单同步接口 APIRequest
alibaba.happytrip.travel.sync

以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
*/
type AlibabaHappytripTravelSyncRequest struct {
    model.Params

    // 差旅申请单对象
    travelHeadDto   *TravelHeadDto 

}

func NewAlibabaHappytripTravelSyncRequest() *AlibabaHappytripTravelSyncRequest{
    return &AlibabaHappytripTravelSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTravelSyncRequest) GetApiMethodName() string {
    return "alibaba.happytrip.travel.sync"
}

func (r AlibabaHappytripTravelSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTravelSyncRequest) SetTravelHeadDto(travelHeadDto *TravelHeadDto) error {
    r.travelHeadDto = travelHeadDto
    r.Set("travel_head_dto", travelHeadDto)
    return nil
}

func (r AlibabaHappytripTravelSyncRequest) GetTravelHeadDto() *TravelHeadDto {
    return r.travelHeadDto
}

