package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalBindshopanditemAPIRequest
绑定门店信息，商品信息 API请求
alibaba.alihealth.reserve.dental.bindshopanditem

绑定门店信息，商品信息 */
type AlibabaAlihealthReserveDentalBindshopanditemAPIRequest struct {
	model.Params
	// bind_list
	_bindList []BindDto
}

// NewAlibabaAlihealthReserveDentalBindshopanditemRequest 初始化AlibabaAlihealthReserveDentalBindshopanditemAPIRequest对象
func NewAlibabaAlihealthReserveDentalBindshopanditemRequest() *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest {
	return &AlibabaAlihealthReserveDentalBindshopanditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.bindshopanditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BindList Setter
// bind_list
func (r *AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) SetBindList(_bindList []BindDto) error {
	r._bindList = _bindList
	r.Set("bind_list", _bindList)
	return nil
}

// Get BindList Getter
func (r AlibabaAlihealthReserveDentalBindshopanditemAPIRequest) GetBindList() []BindDto {
	return r._bindList
}
