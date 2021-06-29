package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取特价菜单 API请求
alibaba.alsc.crm.menu.list

获取特价菜单
*/
type AlibabaAlscCrmMenuListRequest struct {
    model.Params
    // 获取特价菜单请求参数
    _menuOpenReq   *MenuOpenReq
}

// 初始化AlibabaAlscCrmMenuListRequest对象
func NewAlibabaAlscCrmMenuListRequest() *AlibabaAlscCrmMenuListRequest{
    return &AlibabaAlscCrmMenuListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMenuListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.menu.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMenuListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MenuOpenReq Setter
// 获取特价菜单请求参数
func (r *AlibabaAlscCrmMenuListRequest) SetMenuOpenReq(_menuOpenReq *MenuOpenReq) error {
    r._menuOpenReq = _menuOpenReq
    r.Set("menu_open_req", _menuOpenReq)
    return nil
}

// MenuOpenReq Getter
func (r AlibabaAlscCrmMenuListRequest) GetMenuOpenReq() *MenuOpenReq {
    return r._menuOpenReq
}
