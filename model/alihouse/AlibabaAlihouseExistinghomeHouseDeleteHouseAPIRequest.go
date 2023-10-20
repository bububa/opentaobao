package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) Reset() {
	r._house = nil
	r.Params.ToZero()
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

var poolAlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeHouseDeleteHouseRequest()
	},
}

// GetAlibabaAlihouseExistinghomeHouseDeleteHouseRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest
func GetAlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest() *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest {
	return poolAlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest.Get().(*AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest 将 AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest(v *AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest.Put(v)
}
