package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户是否含有会员权益配置菜单权限 APIRequest
alibaba.westcrm.user.benefit.authority

获取指定用户是否含有会员权益配置菜单权限
*/
type AlibabaWestcrmUserBenefitAuthorityRequest struct {
    model.Params

    // 园区ID
    campusId   int64 

    // 当前用户id
    ibUserId   int64 

}

func NewAlibabaWestcrmUserBenefitAuthorityRequest() *AlibabaWestcrmUserBenefitAuthorityRequest{
    return &AlibabaWestcrmUserBenefitAuthorityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetApiMethodName() string {
    return "alibaba.westcrm.user.benefit.authority"
}

func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmUserBenefitAuthorityRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaWestcrmUserBenefitAuthorityRequest) SetIbUserId(ibUserId int64) error {
    r.ibUserId = ibUserId
    r.Set("ib_user_id", ibUserId)
    return nil
}

func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetIbUserId() int64 {
    return r.ibUserId
}

