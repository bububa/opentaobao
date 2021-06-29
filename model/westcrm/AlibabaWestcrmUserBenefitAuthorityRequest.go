package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户是否含有会员权益配置菜单权限 API请求
alibaba.westcrm.user.benefit.authority

获取指定用户是否含有会员权益配置菜单权限
*/
type AlibabaWestcrmUserBenefitAuthorityRequest struct {
    model.Params
    // 园区ID
    _campusId   int64
    // 当前用户id
    _ibUserId   int64
}

// 初始化AlibabaWestcrmUserBenefitAuthorityRequest对象
func NewAlibabaWestcrmUserBenefitAuthorityRequest() *AlibabaWestcrmUserBenefitAuthorityRequest{
    return &AlibabaWestcrmUserBenefitAuthorityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetApiMethodName() string {
    return "alibaba.westcrm.user.benefit.authority"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampusId Setter
// 园区ID
func (r *AlibabaWestcrmUserBenefitAuthorityRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetCampusId() int64 {
    return r._campusId
}
// IbUserId Setter
// 当前用户id
func (r *AlibabaWestcrmUserBenefitAuthorityRequest) SetIbUserId(_ibUserId int64) error {
    r._ibUserId = _ibUserId
    r.Set("ib_user_id", _ibUserId)
    return nil
}

// IbUserId Getter
func (r AlibabaWestcrmUserBenefitAuthorityRequest) GetIbUserId() int64 {
    return r._ibUserId
}
