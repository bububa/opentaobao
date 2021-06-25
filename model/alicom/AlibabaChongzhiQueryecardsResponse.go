package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询指定商家的可用的话费宝贝接口 APIResponse
alibaba.chongzhi.queryecards

查询指定商家的可用的话费宝贝
*/
type AlibabaChongzhiQueryecardsAPIResponse struct {
    model.CommonResponse
    Response *AlibabaChongzhiQueryecardsResponse `json:"alibaba_chongzhi_queryecards_response,omitempty"`
}

type AlibabaChongzhiQueryecardsResponse struct {

    // 结果
    Result   int64 `json:"result,omitempty"`

    // desc
    Desc   string `json:"desc,omitempty"`

    // 订单列表
    EcardList   []EcardItemDo `json:"ecard_list,omitempty"`

    // MtsInfoDo
    CatInfo   *MtsInfoDo `json:"cat_info,omitempty"`

}
