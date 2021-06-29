package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单据后处理状态查询 APIRequest
alibaba.alihealth.drugtrace.top.yljg.query.billstatus

单据处理状态查询
*/
type AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 开始日期
    beginDate   string 

    // 结束日期
    endDate   string 

    // 单据类型 A：全部 AI：全部入库 AO：全部出库
    billType   string 

    // 单据号
    billCode   string 

    // 药品类型
    drugType   string 

    // 状态  0, 上传成功     3, 处理成功     4, 处理失败
    dealStatus   string 

    // 发货商
    fromUserId   string 

    // 收货商
    toUserId   string 

    // 代理商
    agentRefUserId   string 

    // 页大小
    pageSize   int64 

    // 页码
    page   int64 

}

func NewAlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest() *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.billstatus"
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetBeginDate() string {
    return r.beginDate
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetBillType() string {
    return r.billType
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetDrugType(drugType string) error {
    r.drugType = drugType
    r.Set("drug_type", drugType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetDrugType() string {
    return r.drugType
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetDealStatus(dealStatus string) error {
    r.dealStatus = dealStatus
    r.Set("deal_status", dealStatus)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetDealStatus() string {
    return r.dealStatus
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetToUserId(toUserId string) error {
    r.toUserId = toUserId
    r.Set("to_user_id", toUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetToUserId() string {
    return r.toUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetAgentRefUserId(agentRefUserId string) error {
    r.agentRefUserId = agentRefUserId
    r.Set("agent_ref_user_id", agentRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetAgentRefUserId() string {
    return r.agentRefUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryBillstatusRequest) GetPage() int64 {
    return r.page
}

