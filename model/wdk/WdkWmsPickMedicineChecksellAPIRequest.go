package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkwmspickmedicinechecksellAPIRequest 联营商药品柜核销 API请求
// wdk.wms.pick.medicine.checksell
//
// 联营商药品柜核销
type WdkwmspickmedicinechecksellAPIRequest struct {
	model.Params
	// 从二维码扫描出的信息
	_uuid string
	// shopId
	_shopId int64
}

// NewWdkwmspickmedicinechecksellRequest 初始化WdkwmspickmedicinechecksellAPIRequest对象
func NewWdkwmspickmedicinechecksellRequest() *WdkwmspickmedicinechecksellAPIRequest {
	return &WdkwmspickmedicinechecksellAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkwmspickmedicinechecksellAPIRequest) GetApiMethodName() string {
	return "wdk.wms.pick.medicine.checksell"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkwmspickmedicinechecksellAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkwmspickmedicinechecksellAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 从二维码扫描出的信息
func (r *WdkwmspickmedicinechecksellAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkwmspickmedicinechecksellAPIRequest) GetUuid() string {
	return r._uuid
}

// SetShopId is ShopId Setter
// shopId
func (r *WdkwmspickmedicinechecksellAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r WdkwmspickmedicinechecksellAPIRequest) GetShopId() int64 {
	return r._shopId
}
