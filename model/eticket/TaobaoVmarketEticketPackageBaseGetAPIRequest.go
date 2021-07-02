package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketPackageBaseGetAPIRequest 获取包基本信息 API请求
// taobao.vmarket.eticket.package.base.get
//
// 获取包基本信息
type TaobaoVmarketEticketPackageBaseGetAPIRequest struct {
	model.Params
	// 包id
	_packageId int64
}

// NewTaobaoVmarketEticketPackageBaseGetRequest 初始化TaobaoVmarketEticketPackageBaseGetAPIRequest对象
func NewTaobaoVmarketEticketPackageBaseGetRequest() *TaobaoVmarketEticketPackageBaseGetAPIRequest {
	return &TaobaoVmarketEticketPackageBaseGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketPackageBaseGetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.package.base.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketPackageBaseGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPackageId is PackageId Setter
// 包id
func (r *TaobaoVmarketEticketPackageBaseGetAPIRequest) SetPackageId(_packageId int64) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// GetPackageId PackageId Getter
func (r TaobaoVmarketEticketPackageBaseGetAPIRequest) GetPackageId() int64 {
	return r._packageId
}
