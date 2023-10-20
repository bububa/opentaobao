package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest 天猫服务平台商家投诉单服务商申诉接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
type TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest struct {
	model.Params
	// 申诉内容
	_remark string
	// 图片凭证
	_images string
	// 投诉单id
	_id int64
}

// NewTmallservicecenteranomalyrecoursehomedecorationappealRequest 初始化TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest对象
func NewTmallservicecenteranomalyrecoursehomedecorationappealRequest() *TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest {
	return &TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.appeal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 申诉内容
func (r *TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) GetRemark() string {
	return r._remark
}

// SetImages is Images Setter
// 图片凭证
func (r *TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) SetImages(_images string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) GetImages() string {
	return r._images
}

// SetId is Id Setter
// 投诉单id
func (r *TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest) GetId() int64 {
	return r._id
}
