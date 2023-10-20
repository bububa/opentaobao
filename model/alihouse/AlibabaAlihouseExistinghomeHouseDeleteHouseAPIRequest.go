package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest 删除房源 API请求
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
type AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest struct {
	model.Params
	// 房源信息
	_house *DeleteHouseDto
}

// NewAlibabaAlihouseExistinghomeHouseDeleteHouseRequest 初始化AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseDeleteHouseRequest() *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.delete.house"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源信息
func (r *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) SetHouse(_house *DeleteHouseDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) GetHouse() *DeleteHouseDto {
	return r._house
}
