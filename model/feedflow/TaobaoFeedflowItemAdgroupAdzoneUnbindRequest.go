package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元内解绑资源位 APIRequest
taobao.feedflow.item.adgroup.adzone.unbind

信息流单元内解绑资源位
*/
type TaobaoFeedflowItemAdgroupAdzoneUnbindRequest struct {
    model.Params

    // 广告位id
    adzoneIdList   []int64 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemAdgroupAdzoneUnbindRequest() *TaobaoFeedflowItemAdgroupAdzoneUnbindRequest{
    return &TaobaoFeedflowItemAdgroupAdzoneUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.unbind"
}

func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) SetAdzoneIdList(adzoneIdList []int64) error {
    r.adzoneIdList = adzoneIdList
    r.Set("adzone_id_list", adzoneIdList)
    return nil
}

func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetAdzoneIdList() []int64 {
    return r.adzoneIdList
}

func (r *TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemAdgroupAdzoneUnbindRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

