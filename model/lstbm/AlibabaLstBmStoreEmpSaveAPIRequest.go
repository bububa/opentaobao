package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstbmstoreempsaveAPIRequest 保存品牌商自有门店和内部业代之间的关系 API请求
// alibaba.lst.bm.store.emp.save
//
// 保存品牌商自有门店和内部业代之间的关系
type AlibabalstbmstoreempsaveAPIRequest struct {
	model.Params
	// 门店id标识
	_storeId string
	// 员工id标识
	_bmEmpId string
}

// NewAlibabalstbmstoreempsaveRequest 初始化AlibabalstbmstoreempsaveAPIRequest对象
func NewAlibabalstbmstoreempsaveRequest() *AlibabalstbmstoreempsaveAPIRequest {
	return &AlibabalstbmstoreempsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstbmstoreempsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.emp.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstbmstoreempsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstbmstoreempsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店id标识
func (r *AlibabalstbmstoreempsaveAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabalstbmstoreempsaveAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetBmEmpId is BmEmpId Setter
// 员工id标识
func (r *AlibabalstbmstoreempsaveAPIRequest) SetBmEmpId(_bmEmpId string) error {
	r._bmEmpId = _bmEmpId
	r.Set("bm_emp_id", _bmEmpId)
	return nil
}

// GetBmEmpId BmEmpId Getter
func (r AlibabalstbmstoreempsaveAPIRequest) GetBmEmpId() string {
	return r._bmEmpId
}
