package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.change.standard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
