package lstbm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存品牌商自有门店和内部业代之间的关系 API请求
alibaba.lst.bm.store.emp.save

保存品牌商自有门店和内部业代之间的关系
*/
type AlibabaLstBmStoreEmpSaveRequest struct {
    model.Params
    // 门店id标识
    _storeId   string
    // 员工id标识
    _bmEmpId   string
}

// 初始化AlibabaLstBmStoreEmpSaveRequest对象
func NewAlibabaLstBmStoreEmpSaveRequest() *AlibabaLstBmStoreEmpSaveRequest{
    return &AlibabaLstBmStoreEmpSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreEmpSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.bm.store.emp.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreEmpSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店id标识
func (r *AlibabaLstBmStoreEmpSaveRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaLstBmStoreEmpSaveRequest) GetStoreId() string {
    return r._storeId
}
// BmEmpId Setter
// 员工id标识
func (r *AlibabaLstBmStoreEmpSaveRequest) SetBmEmpId(_bmEmpId string) error {
    r._bmEmpId = _bmEmpId
    r.Set("bm_emp_id", _bmEmpId)
    return nil
}

// BmEmpId Getter
func (r AlibabaLstBmStoreEmpSaveRequest) GetBmEmpId() string {
    return r._bmEmpId
}
