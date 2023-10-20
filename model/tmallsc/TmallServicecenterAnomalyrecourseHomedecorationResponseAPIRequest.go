package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest 天猫服务平台商家投诉单服务商响应接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
type TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest struct {
	model.Params
	// 响应方案
	_remark string
	// 图片凭证
	_images string
	// 投诉单id
	_id int64
	// 是否已联系商家；0：未联系；1：已联系
	_contactSeller int64
}

// NewTmallservicecenteranomalyrecoursehomedecorationresponseRequest 初始化TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest对象
func NewTmallservicecenteranomalyrecoursehomedecorationresponseRequest() *TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest {
	return &TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.response"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 响应方案
func (r *TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetRemark() string {
	return r._remark
}

// SetImages is Images Setter
// 图片凭证
func (r *TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) SetImages(_images string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetImages() string {
	return r._images
}

// SetId is Id Setter
// 投诉单id
func (r *TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetId() int64 {
	return r._id
}

// SetContactSeller is ContactSeller Setter
// 是否已联系商家；0：未联系；1：已联系
func (r *TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) SetContactSeller(_contactSeller int64) error {
	r._contactSeller = _contactSeller
	r.Set("contact_seller", _contactSeller)
	return nil
}

// GetContactSeller ContactSeller Getter
func (r TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest) GetContactSeller() int64 {
	return r._contactSeller
}
