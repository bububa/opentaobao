package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询门店 APIResponse
alibaba.lsy.crm.activity.store.getstorelist

ISV查询门店
*/
type AlibabaLsyCrmActivityStoreGetstorelistAPIResponse struct {
    model.CommonResponse
    AlibabaLsyCrmActivityStoreGetstorelistResponse
}

type AlibabaLsyCrmActivityStoreGetstorelistResponse struct {
    XMLName xml.Name `xml:"alibaba_lsy_crm_activity_store_getstorelist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果集
    
    PageResultDO   *PageResultDo `json:"page_result_d_o,omitempty" xml:"page_result_d_o,omitempty"`

    
}
