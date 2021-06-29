package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流删除创意 APIRequest
taobao.feedflow.item.creative.delete

信息流删除创意
*/
type TaobaoFeedflowItemCreativeDeleteRequest struct {
    model.Params

    // 创意id列表
    creativeIdList   []int64 

}

func NewTaobaoFeedflowItemCreativeDeleteRequest() *TaobaoFeedflowItemCreativeDeleteRequest{
    return &TaobaoFeedflowItemCreativeDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemCreativeDeleteRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.creative.delete"
}

func (r TaobaoFeedflowItemCreativeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemCreativeDeleteRequest) SetCreativeIdList(creativeIdList []int64) error {
    r.creativeIdList = creativeIdList
    r.Set("creative_id_list", creativeIdList)
    return nil
}

func (r TaobaoFeedflowItemCreativeDeleteRequest) GetCreativeIdList() []int64 {
    return r.creativeIdList
}

