package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
差旅申请单同步接口 API请求
alibaba.happytrip.travel.sync

以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
*/
type AlibabaHappytripTravelSyncRequest struct {
    model.Params
    // 差旅申请单对象
    travelHeadDto   *TravelHeadDto
}

// 初始化AlibabaHappytripTravelSyncRequest对象
func NewAlibabaHappytripTravelSyncRequest() *AlibabaHappytripTravelSyncRequest{
    return &AlibabaHappytripTravelSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTravelSyncRequest) GetApiMethodName() string {
    return "alibaba.happytrip.travel.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTravelSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TravelHeadDto Setter
// 差旅申请单对象
func (r *AlibabaHappytripTravelSyncRequest) SetTravelHeadDto(travelHeadDto *TravelHeadDto) error {
    r.travelHeadDto = travelHeadDto
    r.Set("travel_head_dto", travelHeadDto)
    return nil
}

// TravelHeadDto Getter
func (r AlibabaHappytripTravelSyncRequest) GetTravelHeadDto() *TravelHeadDto {
    return r.travelHeadDto
}
