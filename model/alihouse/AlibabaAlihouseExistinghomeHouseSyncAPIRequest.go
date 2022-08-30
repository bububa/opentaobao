package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseSyncAPIRequest 房源信息同步 API请求
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
type AlibabaAlihouseExistinghomeHouseSyncAPIRequest struct {
	model.Params
	// 房源信息
	_house *SyncExistingHouseDto
}

// NewAlibabaAlihouseExistinghomeHouseSyncRequest 初始化AlibabaAlihouseExistinghomeHouseSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseSyncRequest() *AlibabaAlihouseExistinghomeHouseSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetHouse is House Setter
// 房源信息
func (r *AlibabaAlihouseExistinghomeHouseSyncAPIRequest) SetHouse(_house *SyncExistingHouseDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaAlihouseExistinghomeHouseSyncAPIRequest) GetHouse() *SyncExistingHouseDto {
	return r._house
}
