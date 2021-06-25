package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户 APIRequest
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系
*/
type AlibabaBaichuanCtgUserRelationRequest struct {
    model.Params

    // 调用的业务方
    app   string 

    // 业务方的用户ID
    uid   string 

    // 淘宝的用户ID
    tbUid   string 

}

func NewAlibabaBaichuanCtgUserRelationRequest() *AlibabaBaichuanCtgUserRelationRequest{
    return &AlibabaBaichuanCtgUserRelationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanCtgUserRelationRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.user.relation"
}

func (r AlibabaBaichuanCtgUserRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBaichuanCtgUserRelationRequest) SetApp(app string) error {
    r.app = app
    r.Set("app", app)
    return nil
}

func (r AlibabaBaichuanCtgUserRelationRequest) GetApp() string {
    return r.app
}

func (r *AlibabaBaichuanCtgUserRelationRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r AlibabaBaichuanCtgUserRelationRequest) GetUid() string {
    return r.uid
}

func (r *AlibabaBaichuanCtgUserRelationRequest) SetTbUid(tbUid string) error {
    r.tbUid = tbUid
    r.Set("tb_uid", tbUid)
    return nil
}

func (r AlibabaBaichuanCtgUserRelationRequest) GetTbUid() string {
    return r.tbUid
}

