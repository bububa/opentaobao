package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawestcrmcustomerinfogetAPIRequest 会员信息查询接口 API请求
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
type AlibabawestcrmcustomerinfogetAPIRequest struct {
	model.Params
	// 支付宝id
	_alipayId string
	// 园区id
	_campusId int64
	// 会员id
	_ibUserId int64
}

// NewAlibabawestcrmcustomerinfogetRequest 初始化AlibabawestcrmcustomerinfogetAPIRequest对象
func NewAlibabawestcrmcustomerinfogetRequest() *AlibabawestcrmcustomerinfogetAPIRequest {
	return &AlibabawestcrmcustomerinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawestcrmcustomerinfogetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.customer.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawestcrmcustomerinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawestcrmcustomerinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayId is AlipayId Setter
// 支付宝id
func (r *AlibabawestcrmcustomerinfogetAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// GetAlipayId AlipayId Getter
func (r AlibabawestcrmcustomerinfogetAPIRequest) GetAlipayId() string {
	return r._alipayId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabawestcrmcustomerinfogetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabawestcrmcustomerinfogetAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetIbUserId is IbUserId Setter
// 会员id
func (r *AlibabawestcrmcustomerinfogetAPIRequest) SetIbUserId(_ibUserId int64) error {
	r._ibUserId = _ibUserId
	r.Set("ib_user_id", _ibUserId)
	return nil
}

// GetIbUserId IbUserId Getter
func (r AlibabawestcrmcustomerinfogetAPIRequest) GetIbUserId() int64 {
	return r._ibUserId
}
