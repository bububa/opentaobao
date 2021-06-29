package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构订单信息 APIRequest
taobao.sebp.organization.getorderinfo

淘小铺合作机构获取机构订单信息，用于机构结算使用
*/
type TaobaoSebpOrganizationGetorderinfoRequest struct {
    model.Params

    // null-请求所有，20200616-请求2020年6月16号的变更信息
    modifyDate   string 

    // 第几页
    pageNum   int64 

    // 查询实时数据时，必传，开始时间结束时间间隔不能超过4个小时
    endTime   string 

    // 查询实时数据时，必传，开始时间不能早于2天前
    startTime   string 

}

func NewTaobaoSebpOrganizationGetorderinfoRequest() *TaobaoSebpOrganizationGetorderinfoRequest{
    return &TaobaoSebpOrganizationGetorderinfoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSebpOrganizationGetorderinfoRequest) GetApiMethodName() string {
    return "taobao.sebp.organization.getorderinfo"
}

func (r TaobaoSebpOrganizationGetorderinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSebpOrganizationGetorderinfoRequest) SetModifyDate(modifyDate string) error {
    r.modifyDate = modifyDate
    r.Set("modify_date", modifyDate)
    return nil
}

func (r TaobaoSebpOrganizationGetorderinfoRequest) GetModifyDate() string {
    return r.modifyDate
}

func (r *TaobaoSebpOrganizationGetorderinfoRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoSebpOrganizationGetorderinfoRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *TaobaoSebpOrganizationGetorderinfoRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSebpOrganizationGetorderinfoRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSebpOrganizationGetorderinfoRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSebpOrganizationGetorderinfoRequest) GetStartTime() string {
    return r.startTime
}

