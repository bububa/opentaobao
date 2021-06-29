package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构上下级关系 API请求
taobao.sebp.organization.getinviteinfo

机构人员获取机构上下级关系信息
*/
type TaobaoSebpOrganizationGetinviteinfoRequest struct {
    model.Params
    // null-请求所有，20200616-请求2020年6月16号的变更信息
    _modifyDate   string
    // 第几页
    _pageNum   int64
}

// 初始化TaobaoSebpOrganizationGetinviteinfoRequest对象
func NewTaobaoSebpOrganizationGetinviteinfoRequest() *TaobaoSebpOrganizationGetinviteinfoRequest{
    return &TaobaoSebpOrganizationGetinviteinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetApiMethodName() string {
    return "taobao.sebp.organization.getinviteinfo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaoSebpOrganizationGetinviteinfoRequest) SetModifyDate(_modifyDate string) error {
    r._modifyDate = _modifyDate
    r.Set("modify_date", _modifyDate)
    return nil
}

// ModifyDate Getter
func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetModifyDate() string {
    return r._modifyDate
}
// PageNum Setter
// 第几页
func (r *TaobaoSebpOrganizationGetinviteinfoRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetPageNum() int64 {
    return r._pageNum
}
