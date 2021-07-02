package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmUserBenefitAuthorityAPIRequest 获取指定用户是否含有会员权益配置菜单权限 API请求
// alibaba.westcrm.user.benefit.authority
//
// 获取指定用户是否含有会员权益配置菜单权限
type AlibabaWestcrmUserBenefitAuthorityAPIRequest struct {
	model.Params
	// 园区ID
	_campusId int64
	// 当前用户id
	_ibUserId int64
}

// NewAlibabaWestcrmUserBenefitAuthorityRequest 初始化AlibabaWestcrmUserBenefitAuthorityAPIRequest对象
func NewAlibabaWestcrmUserBenefitAuthorityRequest() *AlibabaWestcrmUserBenefitAuthorityAPIRequest {
	return &AlibabaWestcrmUserBenefitAuthorityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUserBenefitAuthorityAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.user.benefit.authority"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUserBenefitAuthorityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampusId Setter
// 园区ID
func (r *AlibabaWestcrmUserBenefitAuthorityAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaWestcrmUserBenefitAuthorityAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// Set is IbUserId Setter
// 当前用户id
func (r *AlibabaWestcrmUserBenefitAuthorityAPIRequest) SetIbUserId(_ibUserId int64) error {
	r._ibUserId = _ibUserId
	r.Set("ib_user_id", _ibUserId)
	return nil
}

// Get IbUserId Getter
func (r AlibabaWestcrmUserBenefitAuthorityAPIRequest) GetIbUserId() int64 {
	return r._ibUserId
}
