package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBmStoreEmpSaveAPIRequest 保存品牌商自有门店和内部业代之间的关系 API请求
// alibaba.lst.bm.store.emp.save
//
// 保存品牌商自有门店和内部业代之间的关系
type AlibabaLstBmStoreEmpSaveAPIRequest struct {
	model.Params
	// 门店id标识
	_storeId string
	// 员工id标识
	_bmEmpId string
}

// NewAlibabaLstBmStoreEmpSaveRequest 初始化AlibabaLstBmStoreEmpSaveAPIRequest对象
func NewAlibabaLstBmStoreEmpSaveRequest() *AlibabaLstBmStoreEmpSaveAPIRequest {
	return &AlibabaLstBmStoreEmpSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreEmpSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.emp.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreEmpSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店id标识
func (r *AlibabaLstBmStoreEmpSaveAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaLstBmStoreEmpSaveAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is BmEmpId Setter
// 员工id标识
func (r *AlibabaLstBmStoreEmpSaveAPIRequest) SetBmEmpId(_bmEmpId string) error {
	r._bmEmpId = _bmEmpId
	r.Set("bm_emp_id", _bmEmpId)
	return nil
}

// Get BmEmpId Getter
func (r AlibabaLstBmStoreEmpSaveAPIRequest) GetBmEmpId() string {
	return r._bmEmpId
}
