package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkWmsPickMedicineChecksellAPIRequest) Reset() {
	r._uuid = ""
	r._shopId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineChecksellAPIRequest) GetApiMethodName() string {
	return "wdk.wms.pick.medicine.checksell"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineChecksellAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkWmsPickMedicineChecksellAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 从二维码扫描出的信息
func (r *WdkWmsPickMedicineChecksellAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkWmsPickMedicineChecksellAPIRequest) GetUuid() string {
	return r._uuid
}

// SetShopId is ShopId Setter
// shopId
func (r *WdkWmsPickMedicineChecksellAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r WdkWmsPickMedicineChecksellAPIRequest) GetShopId() int64 {
	return r._shopId
}

var poolWdkWmsPickMedicineChecksellAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkWmsPickMedicineChecksellRequest()
	},
}

// GetWdkWmsPickMedicineChecksellRequest 从 sync.Pool 获取 WdkWmsPickMedicineChecksellAPIRequest
func GetWdkWmsPickMedicineChecksellAPIRequest() *WdkWmsPickMedicineChecksellAPIRequest {
	return poolWdkWmsPickMedicineChecksellAPIRequest.Get().(*WdkWmsPickMedicineChecksellAPIRequest)
}

// ReleaseWdkWmsPickMedicineChecksellAPIRequest 将 WdkWmsPickMedicineChecksellAPIRequest 放入 sync.Pool
func ReleaseWdkWmsPickMedicineChecksellAPIRequest(v *WdkWmsPickMedicineChecksellAPIRequest) {
	v.Reset()
	poolWdkWmsPickMedicineChecksellAPIRequest.Put(v)
}
