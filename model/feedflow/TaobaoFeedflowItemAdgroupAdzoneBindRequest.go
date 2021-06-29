package feedflow

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元内绑定资源位 APIRequest
taobao.feedflow.item.adgroup.adzone.bind

信息流单元内绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzoneBindRequest struct {
    model.Params

    // 新增的绑定资源位
    bindAdzoneList   []AdzoneBindDto 

    // 单元id
    adgroupId   int64 

}

func NewTaobaoFeedflowItemAdgroupAdzoneBindRequest() *TaobaoFeedflowItemAdgroupAdzoneBindRequest{
    return &TaobaoFeedflowItemAdgroupAdzoneBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetApiMethodName() string {
    return "taobao.feedflow.item.adgroup.adzone.bind"
}

func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFeedflowItemAdgroupAdzoneBindRequest) SetBindAdzoneList(bindAdzoneList []AdzoneBindDto) error {
    r.bindAdzoneList = bindAdzoneList
    r.Set("bind_adzone_list", bindAdzoneList)
    return nil
}

func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetBindAdzoneList() []AdzoneBindDto {
    return r.bindAdzoneList
}

func (r *TaobaoFeedflowItemAdgroupAdzoneBindRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoFeedflowItemAdgroupAdzoneBindRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

