package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest 上游出库单单据明细查询 API请求
// alibaba.alihealth.drugtrace.top.lsyd.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 单据编码
	_billCode string
	// 发货企业renEntId
	_fromRefUserId string
	// 收货企业refEntId
	_toRefUserId string
}

// NewAlibabaalihealthdrugtracetoplsydlistupoutdetailRequest 初始化AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest对象
func NewAlibabaalihealthdrugtracetoplsydlistupoutdetailRequest() *AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest {
	return &AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.listupout.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetFromRefUserId is FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// GetFromRefUserId FromRefUserId Getter
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// SetToRefUserId is ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// GetToRefUserId ToRefUserId Getter
func (r AlibabaalihealthdrugtracetoplsydlistupoutdetailAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}
