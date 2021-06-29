package lstbm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存品牌商自有门店和内部业代之间的关系 APIRequest
alibaba.lst.bm.store.emp.save

保存品牌商自有门店和内部业代之间的关系
*/
type AlibabaLstBmStoreEmpSaveRequest struct {
    model.Params

    // 门店id标识
    storeId   string 

    // 员工id标识
    bmEmpId   string 

}

func NewAlibabaLstBmStoreEmpSaveRequest() *AlibabaLstBmStoreEmpSaveRequest{
    return &AlibabaLstBmStoreEmpSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstBmStoreEmpSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.bm.store.emp.save"
}

func (r AlibabaLstBmStoreEmpSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstBmStoreEmpSaveRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaLstBmStoreEmpSaveRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaLstBmStoreEmpSaveRequest) SetBmEmpId(bmEmpId string) error {
    r.bmEmpId = bmEmpId
    r.Set("bm_emp_id", bmEmpId)
    return nil
}

func (r AlibabaLstBmStoreEmpSaveRequest) GetBmEmpId() string {
    return r.bmEmpId
}

