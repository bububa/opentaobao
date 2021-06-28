package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户数娱游戏礼包查询 APIRequest
taobao.deg.user.gamegift.query

查询用户数娱礼包列表
*/
type TaobaoDegUserGamegiftQueryRequest struct {
    model.Params

    // 状态，1为待发放，2为已发放，3为过期
    status   int64 

    // cp item id列表
    cpItemIds   []string 

}

func NewTaobaoDegUserGamegiftQueryRequest() *TaobaoDegUserGamegiftQueryRequest{
    return &TaobaoDegUserGamegiftQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegUserGamegiftQueryRequest) GetApiMethodName() string {
    return "taobao.deg.user.gamegift.query"
}

func (r TaobaoDegUserGamegiftQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegUserGamegiftQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoDegUserGamegiftQueryRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoDegUserGamegiftQueryRequest) SetCpItemIds(cpItemIds []string) error {
    r.cpItemIds = cpItemIds
    r.Set("cp_item_ids", cpItemIds)
    return nil
}

func (r TaobaoDegUserGamegiftQueryRequest) GetCpItemIds() []string {
    return r.cpItemIds
}

