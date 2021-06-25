package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取特价菜单 APIRequest
alibaba.alsc.crm.menu.list

获取特价菜单
*/
type AlibabaAlscCrmMenuListRequest struct {
    model.Params

    // 获取特价菜单请求参数
    menuOpenReq   *MenuOpenReq 

}

func NewAlibabaAlscCrmMenuListRequest() *AlibabaAlscCrmMenuListRequest{
    return &AlibabaAlscCrmMenuListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmMenuListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.menu.list"
}

func (r AlibabaAlscCrmMenuListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmMenuListRequest) SetMenuOpenReq(menuOpenReq *MenuOpenReq) error {
    r.menuOpenReq = menuOpenReq
    r.Set("menu_open_req", menuOpenReq)
    return nil
}

func (r AlibabaAlscCrmMenuListRequest) GetMenuOpenReq() *MenuOpenReq {
    return r.menuOpenReq
}

