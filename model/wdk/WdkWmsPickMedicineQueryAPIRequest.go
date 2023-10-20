package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkwmspickmedicinequeryAPIRequest 查询拣货单中的药品信息 API请求
// wdk.wms.pick.medicine.query
//
// 联营商药机查询拣货单中的药品信息
type WdkwmspickmedicinequeryAPIRequest struct {
	model.Params
	// uuid
	_uuid string
	// shopId
	_shopId int64
}

// NewWdkwmspickmedicinequeryRequest 初始化WdkwmspickmedicinequeryAPIRequest对象
func NewWdkwmspickmedicinequeryRequest() *WdkwmspickmedicinequeryAPIRequest {
	return &WdkwmspickmedicinequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkwmspickmedicinequeryAPIRequest) GetApiMethodName() string {
	return "wdk.wms.pick.medicine.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkwmspickmedicinequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkwmspickmedicinequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// uuid
func (r *WdkwmspickmedicinequeryAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkwmspickmedicinequeryAPIRequest) GetUuid() string {
	return r._uuid
}

// SetShopId is ShopId Setter
// shopId
func (r *WdkwmspickmedicinequeryAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r WdkwmspickmedicinequeryAPIRequest) GetShopId() int64 {
	return r._shopId
}
