package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商场列表 APIResponse
alibaba.westcrm.mall.list.get

根据园区id获取商场列表
*/
type AlibabaWestcrmMallListGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmMallListGetResponse
}

type AlibabaWestcrmMallListGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_mall_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
