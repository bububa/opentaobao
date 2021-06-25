package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商或分销商获取合作关系信息 APIRequest
taobao.fenxiao.cooperation.get

获取供应商的合作关系信息
*/
type TaobaoFenxiaoCooperationGetRequest struct {
    model.Params

    // 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
    status   string 

    // 合作开始时间yyyy-MM-dd HH:mm:ss
    startDate   string 

    // 合作结束时间yyyy-MM-dd HH:mm:ss
    endDate   string 

    // 分销方式：AGENT(代销) 、DEALER（经销）
    tradeType   string 

    // 页码（大于0的整数，默认1）
    pageNo   int64 

    // 每页记录数（默认20，最大50）
    pageSize   int64 

    // 渠道code
    channelCode   string 

    // 1是供应商，2 是分销商
    roleType   int64 

}

func NewTaobaoFenxiaoCooperationGetRequest() *TaobaoFenxiaoCooperationGetRequest{
    return &TaobaoFenxiaoCooperationGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoCooperationGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.cooperation.get"
}

func (r TaobaoFenxiaoCooperationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoCooperationGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetTradeType(tradeType string) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetTradeType() string {
    return r.tradeType
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetChannelCode() string {
    return r.channelCode
}

func (r *TaobaoFenxiaoCooperationGetRequest) SetRoleType(roleType int64) error {
    r.roleType = roleType
    r.Set("role_type", roleType)
    return nil
}

func (r TaobaoFenxiaoCooperationGetRequest) GetRoleType() int64 {
    return r.roleType
}

