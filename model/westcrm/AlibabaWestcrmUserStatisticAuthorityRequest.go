package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户是否含有会员权益统计权限 APIRequest
alibaba.westcrm.user.statistic.authority

获取指定用户是否含有会员权益统计权限
*/
type AlibabaWestcrmUserStatisticAuthorityRequest struct {
    model.Params

    // 园区id
    campusId   int64 

    // 用户id
    ibUserId   int64 

}

func NewAlibabaWestcrmUserStatisticAuthorityRequest() *AlibabaWestcrmUserStatisticAuthorityRequest{
    return &AlibabaWestcrmUserStatisticAuthorityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmUserStatisticAuthorityRequest) GetApiMethodName() string {
    return "alibaba.westcrm.user.statistic.authority"
}

func (r AlibabaWestcrmUserStatisticAuthorityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmUserStatisticAuthorityRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmUserStatisticAuthorityRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaWestcrmUserStatisticAuthorityRequest) SetIbUserId(ibUserId int64) error {
    r.ibUserId = ibUserId
    r.Set("ib_user_id", ibUserId)
    return nil
}

func (r AlibabaWestcrmUserStatisticAuthorityRequest) GetIbUserId() int64 {
    return r.ibUserId
}

