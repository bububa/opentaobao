package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTravelSyncAPIRequest 差旅申请单同步接口 API请求
// alibaba.happytrip.travel.sync
//
// 以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
type AlibabaHappytripTravelSyncAPIRequest struct {
	model.Params
	// 差旅申请单对象
	_travelHeadDto *TravelHeadDto
}

// NewAlibabaHappytripTravelSyncRequest 初始化AlibabaHappytripTravelSyncAPIRequest对象
func NewAlibabaHappytripTravelSyncRequest() *AlibabaHappytripTravelSyncAPIRequest {
	return &AlibabaHappytripTravelSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTravelSyncAPIRequest) Reset() {
	r._travelHeadDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTravelSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.travel.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTravelSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTravelSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTravelHeadDto is TravelHeadDto Setter
// 差旅申请单对象
func (r *AlibabaHappytripTravelSyncAPIRequest) SetTravelHeadDto(_travelHeadDto *TravelHeadDto) error {
	r._travelHeadDto = _travelHeadDto
	r.Set("travel_head_dto", _travelHeadDto)
	return nil
}

// GetTravelHeadDto TravelHeadDto Getter
func (r AlibabaHappytripTravelSyncAPIRequest) GetTravelHeadDto() *TravelHeadDto {
	return r._travelHeadDto
}

var poolAlibabaHappytripTravelSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTravelSyncRequest()
	},
}

// GetAlibabaHappytripTravelSyncRequest 从 sync.Pool 获取 AlibabaHappytripTravelSyncAPIRequest
func GetAlibabaHappytripTravelSyncAPIRequest() *AlibabaHappytripTravelSyncAPIRequest {
	return poolAlibabaHappytripTravelSyncAPIRequest.Get().(*AlibabaHappytripTravelSyncAPIRequest)
}

// ReleaseAlibabaHappytripTravelSyncAPIRequest 将 AlibabaHappytripTravelSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTravelSyncAPIRequest(v *AlibabaHappytripTravelSyncAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTravelSyncAPIRequest.Put(v)
}
