package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店亲橙客栈预订信息同步 API请求
alibaba.htorder.hotel.sync.booking

未来酒店亲橙客栈预订信息同步
*/
type AlibabaHtorderHotelSyncBookingRequest struct {
    model.Params
    // 预订信息数据
    _dataEntity   *SyncHotelBookingDataRequestDto
}

// 初始化AlibabaHtorderHotelSyncBookingRequest对象
func NewAlibabaHtorderHotelSyncBookingRequest() *AlibabaHtorderHotelSyncBookingRequest{
    return &AlibabaHtorderHotelSyncBookingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncBookingRequest) GetApiMethodName() string {
    return "alibaba.htorder.hotel.sync.booking"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncBookingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataEntity Setter
// 预订信息数据
func (r *AlibabaHtorderHotelSyncBookingRequest) SetDataEntity(_dataEntity *SyncHotelBookingDataRequestDto) error {
    r._dataEntity = _dataEntity
    r.Set("data_entity", _dataEntity)
    return nil
}

// DataEntity Getter
func (r AlibabaHtorderHotelSyncBookingRequest) GetDataEntity() *SyncHotelBookingDataRequestDto {
    return r._dataEntity
}
