package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest 天猫服务平台商家投诉单服务商完结接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.close
//
// 天猫服务平台商家投诉单服务商完结接口
type TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest struct {
	model.Params
	// 完结方案
	_remark string
	// 图片凭证
	_images string
	// 投诉单id
	_id int64
	// 是否与商家达成一致；0：未达成；1：已达成
	_isAgreement int64
}

// NewTmallservicecenteranomalyrecoursehomedecorationcloseRequest 初始化TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest对象
func NewTmallservicecenteranomalyrecoursehomedecorationcloseRequest() *TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest {
	return &TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 完结方案
func (r *TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetRemark() string {
	return r._remark
}

// SetImages is Images Setter
// 图片凭证
func (r *TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) SetImages(_images string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetImages() string {
	return r._images
}

// SetId is Id Setter
// 投诉单id
func (r *TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetId() int64 {
	return r._id
}

// SetIsAgreement is IsAgreement Setter
// 是否与商家达成一致；0：未达成；1：已达成
func (r *TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) SetIsAgreement(_isAgreement int64) error {
	r._isAgreement = _isAgreement
	r.Set("is_agreement", _isAgreement)
	return nil
}

// GetIsAgreement IsAgreement Getter
func (r TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest) GetIsAgreement() int64 {
	return r._isAgreement
}
