package charity

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询公益3小时公益时汇总 API请求
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

// 初始化AlibabaCharityCharitytimeQueryRequest对象
func NewAlibabaCharityCharitytimeQueryRequest() *AlibabaCharityCharitytimeQueryRequest{
    return &AlibabaCharityCharitytimeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeQueryRequest) GetApiMethodName() string {
    return "alibaba.charity.charitytime.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CharityTypeList Setter
// 公益类型
func (r *AlibabaCharityCharitytimeQueryRequest) SetCharityTypeList(charityTypeList []string) error {
    r.charityTypeList = charityTypeList
    r.Set("charity_type_list", charityTypeList)
    return nil
}

// CharityTypeList Getter
func (r AlibabaCharityCharitytimeQueryRequest) GetCharityTypeList() []string {
    return r.charityTypeList
}
// EndDate Setter
// 结束时间戳-毫秒时间
func (r *AlibabaCharityCharitytimeQueryRequest) SetEndDate(endDate int64) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaCharityCharitytimeQueryRequest) GetEndDate() int64 {
    return r.endDate
}
// StartDate Setter
// 开始时间戳-毫秒时间
func (r *AlibabaCharityCharitytimeQueryRequest) SetStartDate(startDate int64) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r AlibabaCharityCharitytimeQueryRequest) GetStartDate() int64 {
    return r.startDate
}
// TbUid Setter
// 淘宝Uid
func (r *AlibabaCharityCharitytimeQueryRequest) SetTbUid(tbUid int64) error {
    r.tbUid = tbUid
    r.Set("tb_uid", tbUid)
    return nil
}

// TbUid Getter
func (r AlibabaCharityCharitytimeQueryRequest) GetTbUid() int64 {
    return r.tbUid
}
// ActivityId Setter
// 活动ID
func (r *AlibabaCharityCharitytimeQueryRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaCharityCharitytimeQueryRequest) GetActivityId() int64 {
    return r.activityId
}
// ExtParam Setter
// 扩展参数
func (r *AlibabaCharityCharitytimeQueryRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

// ExtParam Getter
func (r AlibabaCharityCharitytimeQueryRequest) GetExtParam() string {
    return r.extParam
}
