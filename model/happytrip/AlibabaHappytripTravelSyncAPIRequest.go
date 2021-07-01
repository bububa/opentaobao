package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTravelSyncAPIRequest
差旅申请单同步接口 API请求
alibaba.happytrip.travel.sync

以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中 */
type AlibabaHappytripTravelSyncAPIRequest struct {
	model.Params
	// 差旅申请单对象
	_travelHeadDto *TravelHeadDto
}

// NewAlibabaHappytripTravelSyncRequest 初始化AlibabaHappytripTravelSyncAPIRequest对象
func NewAlibabaHappytripTravelSyncRequest() *AlibabaHappytripTravelSyncAPIRequest {
	return &AlibabaHappytripTravelSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTravelSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.travel.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTravelSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TravelHeadDto Setter
// 差旅申请单对象
func (r *AlibabaHappytripTravelSyncAPIRequest) SetTravelHeadDto(_travelHeadDto *TravelHeadDto) error {
	r._travelHeadDto = _travelHeadDto
	r.Set("travel_head_dto", _travelHeadDto)
	return nil
}

// Get TravelHeadDto Getter
func (r AlibabaHappytripTravelSyncAPIRequest) GetTravelHeadDto() *TravelHeadDto {
	return r._travelHeadDto
}
