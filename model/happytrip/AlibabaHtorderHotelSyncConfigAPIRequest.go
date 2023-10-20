package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahtorderhotelsyncconfigAPIRequest 同步配置信息 API请求
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
type AlibabahtorderhotelsyncconfigAPIRequest struct {
	model.Params
	// 配置信息
	_dataEntity *HotelMessageConfigDto
}

// NewAlibabahtorderhotelsyncconfigRequest 初始化AlibabahtorderhotelsyncconfigAPIRequest对象
func NewAlibabahtorderhotelsyncconfigRequest() *AlibabahtorderhotelsyncconfigAPIRequest {
	return &AlibabahtorderhotelsyncconfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahtorderhotelsyncconfigAPIRequest) GetApiMethodName() string {
	return "alibaba.htorder.hotel.sync.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahtorderhotelsyncconfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahtorderhotelsyncconfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataEntity is DataEntity Setter
// 配置信息
func (r *AlibabahtorderhotelsyncconfigAPIRequest) SetDataEntity(_dataEntity *HotelMessageConfigDto) error {
	r._dataEntity = _dataEntity
	r.Set("data_entity", _dataEntity)
	return nil
}

// GetDataEntity DataEntity Getter
func (r AlibabahtorderhotelsyncconfigAPIRequest) GetDataEntity() *HotelMessageConfigDto {
	return r._dataEntity
}
