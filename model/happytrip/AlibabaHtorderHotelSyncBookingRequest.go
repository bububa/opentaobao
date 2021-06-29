package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店亲橙客栈预订信息同步 APIRequest
alibaba.htorder.hotel.sync.booking

未来酒店亲橙客栈预订信息同步
*/
type AlibabaHtorderHotelSyncBookingRequest struct {
    model.Params

    // 预订信息数据
    dataEntity   *SyncHotelBookingDataRequestDto 

}

func NewAlibabaHtorderHotelSyncBookingRequest() *AlibabaHtorderHotelSyncBookingRequest{
    return &AlibabaHtorderHotelSyncBookingRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHtorderHotelSyncBookingRequest) GetApiMethodName() string {
    return "alibaba.htorder.hotel.sync.booking"
}

func (r AlibabaHtorderHotelSyncBookingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHtorderHotelSyncBookingRequest) SetDataEntity(dataEntity *SyncHotelBookingDataRequestDto) error {
    r.dataEntity = dataEntity
    r.Set("data_entity", dataEntity)
    return nil
}

func (r AlibabaHtorderHotelSyncBookingRequest) GetDataEntity() *SyncHotelBookingDataRequestDto {
    return r.dataEntity
}

