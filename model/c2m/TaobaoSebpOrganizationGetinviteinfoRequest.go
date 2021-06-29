package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构上下级关系 APIRequest
taobao.sebp.organization.getinviteinfo

机构人员获取机构上下级关系信息
*/
type TaobaoSebpOrganizationGetinviteinfoRequest struct {
    model.Params

    // null-请求所有，20200616-请求2020年6月16号的变更信息
    modifyDate   string 

    // 第几页
    pageNum   int64 

}

func NewTaobaoSebpOrganizationGetinviteinfoRequest() *TaobaoSebpOrganizationGetinviteinfoRequest{
    return &TaobaoSebpOrganizationGetinviteinfoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetApiMethodName() string {
    return "taobao.sebp.organization.getinviteinfo"
}

func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSebpOrganizationGetinviteinfoRequest) SetModifyDate(modifyDate string) error {
    r.modifyDate = modifyDate
    r.Set("modify_date", modifyDate)
    return nil
}

func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetModifyDate() string {
    return r.modifyDate
}

func (r *TaobaoSebpOrganizationGetinviteinfoRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoSebpOrganizationGetinviteinfoRequest) GetPageNum() int64 {
    return r.pageNum
}

