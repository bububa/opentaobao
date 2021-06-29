package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的话费宝贝接口 APIResponse
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝
*/
type AlibabaChongzhiQueryecardsAPIResponse struct {
    model.CommonResponse
    AlibabaChongzhiQueryecardsResponse
}

type AlibabaChongzhiQueryecardsResponse struct {
    XMLName xml.Name `xml:"alibaba_chongzhi_queryecards_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
    // desc
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`

    
    // 订单列表
    
    EcardList   []EcardItemDo `json:"ecard_list,omitempty" xml:"ecard_list>ecard_item_do,omitempty"`
    
    
    // MtsInfoDo
    
    CatInfo   *MtsInfoDo `json:"cat_info,omitempty" xml:"cat_info,omitempty"`

    
}
