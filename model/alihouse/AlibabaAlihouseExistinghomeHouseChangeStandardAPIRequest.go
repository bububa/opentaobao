package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest 委托房源变更标准房源 API请求
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
type AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest struct {
	model.Params
	// 房源信息
	_house *EntrustChangeStandardDto
}

// NewAlibabaAlihouseExistinghomeHouseChangeStandardRequest 初始化AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseChangeStandardRequest() *AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) Reset() {
	r._house = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.change.standard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源信息
func (r *AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) SetHouse(_house *EntrustChangeStandardDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) GetHouse() *EntrustChangeStandardDto {
	return r._house
}

var poolAlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeHouseChangeStandardRequest()
	},
}

// GetAlibabaAlihouseExistinghomeHouseChangeStandardRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest
func GetAlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest() *AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest {
	return poolAlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest.Get().(*AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest 将 AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest(v *AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest.Put(v)
}
