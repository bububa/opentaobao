package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest 天猫服务平台商家投诉单服务商认责接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.admit
//
// 天猫服务平台商家投诉单服务商认责接口
type TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest struct {
	model.Params
	// 备注
	_remark string
	// 投诉单id
	_id int64
	// 认责金额，分
	_tpAdmitResponsibleAmount int64
}

// NewTmallservicecenteranomalyrecoursehomedecorationadmitRequest 初始化TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest对象
func NewTmallservicecenteranomalyrecoursehomedecorationadmitRequest() *TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest {
	return &TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.admit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 备注
func (r *TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) GetRemark() string {
	return r._remark
}

// SetId is Id Setter
// 投诉单id
func (r *TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) GetId() int64 {
	return r._id
}

// SetTpAdmitResponsibleAmount is TpAdmitResponsibleAmount Setter
// 认责金额，分
func (r *TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) SetTpAdmitResponsibleAmount(_tpAdmitResponsibleAmount int64) error {
	r._tpAdmitResponsibleAmount = _tpAdmitResponsibleAmount
	r.Set("tp_admit_responsible_amount", _tpAdmitResponsibleAmount)
	return nil
}

// GetTpAdmitResponsibleAmount TpAdmitResponsibleAmount Getter
func (r TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest) GetTpAdmitResponsibleAmount() int64 {
	return r._tpAdmitResponsibleAmount
}
