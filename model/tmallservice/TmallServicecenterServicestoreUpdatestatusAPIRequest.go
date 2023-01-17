package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreUpdatestatusAPIRequest 网点/门店状态修改 API请求
// tmall.servicecenter.servicestore.updatestatus
//
// 修改网点/门店状态
type TmallServicecenterServicestoreUpdatestatusAPIRequest struct {
	model.Params
	// 业务类型。不同业务传不同的值
	_bizType string
	// 门店id
	_id int64
	// 状态。1 营业，0歇业，-1彻底关店
	_status int64
}

// NewTmallServicecenterServicestoreUpdatestatusRequest 初始化TmallServicecenterServicestoreUpdatestatusAPIRequest对象
func NewTmallServicecenterServicestoreUpdatestatusRequest() *TmallServicecenterServicestoreUpdatestatusAPIRequest {
	return &TmallServicecenterServicestoreUpdatestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdatestatusAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicestore.updatestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdatestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicestoreUpdatestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务类型。不同业务传不同的值
func (r *TmallServicecenterServicestoreUpdatestatusAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallServicecenterServicestoreUpdatestatusAPIRequest) GetBizType() string {
	return r._bizType
}

// SetId is Id Setter
// 门店id
func (r *TmallServicecenterServicestoreUpdatestatusAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterServicestoreUpdatestatusAPIRequest) GetId() int64 {
	return r._id
}

// SetStatus is Status Setter
// 状态。1 营业，0歇业，-1彻底关店
func (r *TmallServicecenterServicestoreUpdatestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallServicecenterServicestoreUpdatestatusAPIRequest) GetStatus() int64 {
	return r._status
}
