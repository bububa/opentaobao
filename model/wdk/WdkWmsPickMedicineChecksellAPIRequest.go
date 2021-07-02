package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkWmsPickMedicineChecksellAPIRequest 联营商药品柜核销 API请求
// wdk.wms.pick.medicine.checksell
//
// 联营商药品柜核销
type WdkWmsPickMedicineChecksellAPIRequest struct {
	model.Params
	// 从二维码扫描出的信息
	_uuid string
	// shopId
	_shopId int64
}

// NewWdkWmsPickMedicineChecksellRequest 初始化WdkWmsPickMedicineChecksellAPIRequest对象
func NewWdkWmsPickMedicineChecksellRequest() *WdkWmsPickMedicineChecksellAPIRequest {
	return &WdkWmsPickMedicineChecksellAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineChecksellAPIRequest) GetApiMethodName() string {
	return "wdk.wms.pick.medicine.checksell"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineChecksellAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Uuid Setter
// 从二维码扫描出的信息
func (r *WdkWmsPickMedicineChecksellAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r WdkWmsPickMedicineChecksellAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is ShopId Setter
// shopId
func (r *WdkWmsPickMedicineChecksellAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r WdkWmsPickMedicineChecksellAPIRequest) GetShopId() int64 {
	return r._shopId
}
