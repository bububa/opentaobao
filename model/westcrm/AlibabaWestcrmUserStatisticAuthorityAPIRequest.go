package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmUserStatisticAuthorityAPIRequest
获取指定用户是否含有会员权益统计权限 API请求
alibaba.westcrm.user.statistic.authority

获取指定用户是否含有会员权益统计权限 */
type AlibabaWestcrmUserStatisticAuthorityAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 用户id
	_ibUserId int64
}

// NewAlibabaWestcrmUserStatisticAuthorityRequest 初始化AlibabaWestcrmUserStatisticAuthorityAPIRequest对象
func NewAlibabaWestcrmUserStatisticAuthorityRequest() *AlibabaWestcrmUserStatisticAuthorityAPIRequest {
	return &AlibabaWestcrmUserStatisticAuthorityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUserStatisticAuthorityAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.user.statistic.authority"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUserStatisticAuthorityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampusId Setter
// 园区id
func (r *AlibabaWestcrmUserStatisticAuthorityAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaWestcrmUserStatisticAuthorityAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// Set is IbUserId Setter
// 用户id
func (r *AlibabaWestcrmUserStatisticAuthorityAPIRequest) SetIbUserId(_ibUserId int64) error {
	r._ibUserId = _ibUserId
	r.Set("ib_user_id", _ibUserId)
	return nil
}

// Get IbUserId Getter
func (r AlibabaWestcrmUserStatisticAuthorityAPIRequest) GetIbUserId() int64 {
	return r._ibUserId
}
