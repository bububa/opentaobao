package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghouseHouseBaseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghouse.house.base"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghouseHouseBaseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
