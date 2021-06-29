package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡模板列表(支持数据下行) APIResponse
alibaba.alsc.crm.card.pagetmp

查询卡模板列表(支持数据下行)
当传递了数据下行参数:
     * isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
     * 否则分页查询卡模板,返回结果带有分页信息
*/
type AlibabaAlscCrmCardPagetmpAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardPagetmpResponse
}

type AlibabaAlscCrmCardPagetmpResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_pagetmp_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
