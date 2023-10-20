package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkWmsPickMedicineQueryAPIRequest 查询拣货单中的药品信息 API请求
// wdk.wms.pick.medicine.query
//
// 联营商药机查询拣货单中的药品信息
type WdkWmsPickMedicineQueryAPIRequest struct {
	model.Params
	// uuid
	_uuid string
	// shopId
	_shopId int64
}

// NewWdkWmsPickMedicineQueryRequest 初始化WdkWmsPickMedicineQueryAPIRequest对象
func NewWdkWmsPickMedicineQueryRequest() *WdkWmsPickMedicineQueryAPIRequest {
	return &WdkWmsPickMedicineQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkWmsPickMedicineQueryAPIRequest) Reset() {
	r._uuid = ""
	r._shopId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkWmsPickMedicineQueryAPIRequest) GetApiMethodName() string {
	return "wdk.wms.pick.medicine.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkWmsPickMedicineQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkWmsPickMedicineQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// uuid
func (r *WdkWmsPickMedicineQueryAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkWmsPickMedicineQueryAPIRequest) GetUuid() string {
	return r._uuid
}

// SetShopId is ShopId Setter
// shopId
func (r *WdkWmsPickMedicineQueryAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r WdkWmsPickMedicineQueryAPIRequest) GetShopId() int64 {
	return r._shopId
}

var poolWdkWmsPickMedicineQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkWmsPickMedicineQueryRequest()
	},
}

// GetWdkWmsPickMedicineQueryRequest 从 sync.Pool 获取 WdkWmsPickMedicineQueryAPIRequest
func GetWdkWmsPickMedicineQueryAPIRequest() *WdkWmsPickMedicineQueryAPIRequest {
	return poolWdkWmsPickMedicineQueryAPIRequest.Get().(*WdkWmsPickMedicineQueryAPIRequest)
}

// ReleaseWdkWmsPickMedicineQueryAPIRequest 将 WdkWmsPickMedicineQueryAPIRequest 放入 sync.Pool
func ReleaseWdkWmsPickMedicineQueryAPIRequest(v *WdkWmsPickMedicineQueryAPIRequest) {
	v.Reset()
	poolWdkWmsPickMedicineQueryAPIRequest.Put(v)
}
