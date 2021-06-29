package lstbm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改品牌商自有门店数据 APIResponse
alibaba.lst.bm.store.update

修改品牌商自有门店数据
*/
type AlibabaLstBmStoreUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLstBmStoreUpdateResponse
}

type AlibabaLstBmStoreUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_bm_store_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true表示执行成功，false表示执行失败
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
