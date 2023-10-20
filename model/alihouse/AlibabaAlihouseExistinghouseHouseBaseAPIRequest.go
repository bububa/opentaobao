package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghouseHouseBaseAPIRequest （租房）同步房屋信息 API请求
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
type AlibabaAlihouseExistinghouseHouseBaseAPIRequest struct {
	model.Params
	// 基础房屋信息
	_houseBase *HouseBaseDto
}

// NewAlibabaAlihouseExistinghouseHouseBaseRequest 初始化AlibabaAlihouseExistinghouseHouseBaseAPIRequest对象
func NewAlibabaAlihouseExistinghouseHouseBaseRequest() *AlibabaAlihouseExistinghouseHouseBaseAPIRequest {
	return &AlibabaAlihouseExistinghouseHouseBaseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghouseHouseBaseAPIRequest) Reset() {
	r._houseBase = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghouseHouseBaseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghouse.house.base"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghouseHouseBaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghouseHouseBaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouseBase is HouseBase Setter
// 基础房屋信息
func (r *AlibabaAlihouseExistinghouseHouseBaseAPIRequest) SetHouseBase(_houseBase *HouseBaseDto) error {
	r._houseBase = _houseBase
	r.Set("house_base", _houseBase)
	return nil
}

// GetHouseBase HouseBase Getter
func (r AlibabaAlihouseExistinghouseHouseBaseAPIRequest) GetHouseBase() *HouseBaseDto {
	return r._houseBase
}

var poolAlibabaAlihouseExistinghouseHouseBaseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghouseHouseBaseRequest()
	},
}

// GetAlibabaAlihouseExistinghouseHouseBaseRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghouseHouseBaseAPIRequest
func GetAlibabaAlihouseExistinghouseHouseBaseAPIRequest() *AlibabaAlihouseExistinghouseHouseBaseAPIRequest {
	return poolAlibabaAlihouseExistinghouseHouseBaseAPIRequest.Get().(*AlibabaAlihouseExistinghouseHouseBaseAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghouseHouseBaseAPIRequest 将 AlibabaAlihouseExistinghouseHouseBaseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghouseHouseBaseAPIRequest(v *AlibabaAlihouseExistinghouseHouseBaseAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghouseHouseBaseAPIRequest.Put(v)
}
