package lstbm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存品牌商自有门店和内部业代之间的关系 APIResponse
alibaba.lst.bm.store.emp.save

保存品牌商自有门店和内部业代之间的关系
*/
type AlibabaLstBmStoreEmpSaveAPIResponse struct {
    model.CommonResponse
    AlibabaLstBmStoreEmpSaveResponse
}

type AlibabaLstBmStoreEmpSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_bm_store_emp_save_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true表示执行成功，false表示执行失败
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
