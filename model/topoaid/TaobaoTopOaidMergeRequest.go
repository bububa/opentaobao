package topoaid

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OAID订单合并 APIRequest
taobao.top.oaid.merge

基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
*/
type TaobaoTopOaidMergeRequest struct {
    model.Params

    // 合单请求列表，最多支持100个。
    mergeList   []OrderMerge 

}

func NewTaobaoTopOaidMergeRequest() *TaobaoTopOaidMergeRequest{
    return &TaobaoTopOaidMergeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopOaidMergeRequest) GetApiMethodName() string {
    return "taobao.top.oaid.merge"
}

func (r TaobaoTopOaidMergeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopOaidMergeRequest) SetMergeList(mergeList []OrderMerge) error {
    r.mergeList = mergeList
    r.Set("merge_list", mergeList)
    return nil
}

func (r TaobaoTopOaidMergeRequest) GetMergeList() []OrderMerge {
    return r.mergeList
}

