package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest 天猫服务平台一键求助单服务商备注更新接口 API请求
// tmall.servicecenter.anomalyrecourse.remark.update
//
// 一键求助服务商可以回传备注
type TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest struct {
	model.Params
	// 需要更新的一键求助单id
	_id int64
	// 需要更新的服务商备注
	_remark string
}

// NewTmallServicecenterAnomalyrecourseRemarkUpdateRequest 初始化TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest对象
func NewTmallServicecenterAnomalyrecourseRemarkUpdateRequest() *TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest {
	return &TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.remark.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 需要更新的一键求助单id
func (r *TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest) GetId() int64 {
	return r._id
}

// SetRemark is Remark Setter
// 需要更新的服务商备注
func (r *TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest) GetRemark() string {
	return r._remark
}
