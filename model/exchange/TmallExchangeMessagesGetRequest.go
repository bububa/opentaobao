package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询换货订单留言列表 APIRequest
tmall.exchange.messages.get

查询换货订单留言列表
*/
type TmallExchangeMessagesGetRequest struct {
    model.Params

    // 留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6)
    operatorRoles   []int64 

    // 每页条数
    pageSize   int64 

    // 换货单号ID
    disputeId   int64 

    // 页码
    pageNo   int64 

    // 返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
    fields   []string 

}

func NewTmallExchangeMessagesGetRequest() *TmallExchangeMessagesGetRequest{
    return &TmallExchangeMessagesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeMessagesGetRequest) GetApiMethodName() string {
    return "tmall.exchange.messages.get"
}

func (r TmallExchangeMessagesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeMessagesGetRequest) SetOperatorRoles(operatorRoles []int64) error {
    r.operatorRoles = operatorRoles
    r.Set("operator_roles", operatorRoles)
    return nil
}

func (r TmallExchangeMessagesGetRequest) GetOperatorRoles() []int64 {
    return r.operatorRoles
}

func (r *TmallExchangeMessagesGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TmallExchangeMessagesGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TmallExchangeMessagesGetRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeMessagesGetRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeMessagesGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TmallExchangeMessagesGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TmallExchangeMessagesGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeMessagesGetRequest) GetFields() []string {
    return r.fields
}

