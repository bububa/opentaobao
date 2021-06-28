package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的流量宝贝接口 APIResponse
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝
*/
type AlibabaChongzhiQueryflowAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaChongzhiQueryflowResponse `json:"alibaba_chongzhi_queryflow_response,omitempty"` 
    AlibabaChongzhiQueryflowResponse
}

/* model for simplify = false
type AlibabaChongzhiQueryflowResponse struct {

    // 结果
    
    Result   int64 `json:"result,omitempty"`
    

    // desc
    
    Desc   string `json:"desc,omitempty"`
    

    // MtsInfoDo
    
    CatInfo  *struct {
        MtsInfoDo  *MtsInfoDo `json:"mts_info_do,omitempty"`
    } `json:"cat_info,omitempty"`
    

    // flow_card_list
    
    FlowCardList  struct {
        Flowcardlist  []Flowcardlist `json:"flowcardlist,omitempty"`
    } `json:"flow_card_list,omitempty"`
    

    // cn_desc
    
    CnDesc   string `json:"cn_desc,omitempty"`
    

}
*/

type AlibabaChongzhiQueryflowResponse struct {

    // 结果
    Result   int64 `json:"result,omitempty"`

    // desc
    Desc   string `json:"desc,omitempty"`

    // MtsInfoDo
    CatInfo   *MtsInfoDo `json:"cat_info,omitempty"`

    // flow_card_list
    FlowCardList   []Flowcardlist `json:"flow_card_list,omitempty"`

    // cn_desc
    CnDesc   string `json:"cn_desc,omitempty"`

}
