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
type AlibabaHtorderHotelSyncBookingAPIRequest struct {
    model.Params
    // 预订信息数据
    _dataEntity   *SyncHotelBookingDataRequestDTO
}

// 初始化AlibabaHtorderHotelSyncBookingAPIRequest对象
func NewAlibabaHtorderHotelSyncBookingRequest() *AlibabaHtorderHotelSyncBookingAPIRequest{
    return &AlibabaHtorderHotelSyncBookingAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetApiMethodName() string {
    return "alibaba.htorder.hotel.sync.booking"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataEntity Setter
// 预订信息数据
func (r *AlibabaHtorderHotelSyncBookingAPIRequest) SetDataEntity(_dataEntity *SyncHotelBookingDataRequestDTO) error {
    r._dataEntity = _dataEntity
    r.Set("data_entity", _dataEntity)
    return nil
}

// DataEntity Getter
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetDataEntity() *SyncHotelBookingDataRequestDTO {
    return r._dataEntity
}
