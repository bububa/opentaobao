package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的流量宝贝接口 APIResponse
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝
*/
type AlibabaChongzhiQueryflowAPIResponse struct {
    model.CommonResponse
    AlibabaChongzhiQueryflowResponse
}

type AlibabaChongzhiQueryflowResponse struct {
    XMLName xml.Name `xml:"alibaba_chongzhi_queryflow_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
    // desc
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`

    
    // MtsInfoDo
    
    CatInfo   *MtsInfoDo `json:"cat_info,omitempty" xml:"cat_info,omitempty"`

    
    // flow_card_list
    
    FlowCardList   []Flowcardlist `json:"flow_card_list,omitempty" xml:"flow_card_list>flowcardlist,omitempty"`
    
    
    // cn_desc
    
    CnDesc   string `json:"cn_desc,omitempty" xml:"cn_desc,omitempty"`

    
}
