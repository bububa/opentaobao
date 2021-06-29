package charity

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询公益3小时公益时汇总 APIRequest
alibaba.charity.charitytime.query

查询公益3小时公益时汇总
*/
type AlibabaCharityCharitytimeQueryRequest struct {
    model.Params

    // 公益类型
    charityTypeList   []string 

    // 结束时间戳-毫秒时间
    endDate   int64 

    // 开始时间戳-毫秒时间
    startDate   int64 

    // 淘宝Uid
    tbUid   int64 

    // 活动ID
    activityId   int64 

    // 扩展参数
    extParam   string 

}

func NewAlibabaCharityCharitytimeQueryRequest() *AlibabaCharityCharitytimeQueryRequest{
    return &AlibabaCharityCharitytimeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCharityCharitytimeQueryRequest) GetApiMethodName() string {
    return "alibaba.charity.charitytime.query"
}

func (r AlibabaCharityCharitytimeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCharityCharitytimeQueryRequest) SetCharityTypeList(charityTypeList []string) error {
    r.charityTypeList = charityTypeList
    r.Set("charity_type_list", charityTypeList)
    return nil
}

func (r AlibabaCharityCharitytimeQueryRequest) GetCharityTypeList() []string {
    return r.charityTypeList
}

func (r *AlibabaCharityCharitytimeQueryRequest) SetEndDate(endDate int64) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaCharityCharitytimeQueryRequest) GetEndDate() int64 {
    return r.endDate
}

func (r *AlibabaCharityCharitytimeQueryRequest) SetStartDate(startDate int64) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r AlibabaCharityCharitytimeQueryRequest) GetStartDate() int64 {
    return r.startDate
}

func (r *AlibabaCharityCharitytimeQueryRequest) SetTbUid(tbUid int64) error {
    r.tbUid = tbUid
    r.Set("tb_uid", tbUid)
    return nil
}

func (r AlibabaCharityCharitytimeQueryRequest) GetTbUid() int64 {
    return r.tbUid
}

func (r *AlibabaCharityCharitytimeQueryRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaCharityCharitytimeQueryRequest) GetActivityId() int64 {
    return r.activityId
}

func (r *AlibabaCharityCharitytimeQueryRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

func (r AlibabaCharityCharitytimeQueryRequest) GetExtParam() string {
    return r.extParam
}

