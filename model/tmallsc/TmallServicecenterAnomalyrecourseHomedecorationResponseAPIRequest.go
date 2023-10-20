package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest 天猫服务平台商家投诉单服务商响应接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
type TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest struct {
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

// NewTmallServicecenterAnomalyrecourseHomedecorationResponseRequest 初始化TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest对象
func NewTmallServicecenterAnomalyrecourseHomedecorationResponseRequest() *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest {
	return &TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) Reset() {
	r._remark = ""
	r._images = ""
	r._id = 0
	r._contactSeller = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.response"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 响应方案
func (r *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetRemark() string {
	return r._remark
}

// SetImages is Images Setter
// 图片凭证
func (r *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) SetImages(_images string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetImages() string {
	return r._images
}

// SetId is Id Setter
// 投诉单id
func (r *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetId() int64 {
	return r._id
}

// SetContactSeller is ContactSeller Setter
// 是否已联系商家；0：未联系；1：已联系
func (r *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) SetContactSeller(_contactSeller int64) error {
	r._contactSeller = _contactSeller
	r.Set("contact_seller", _contactSeller)
	return nil
}

// GetContactSeller ContactSeller Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) GetContactSeller() int64 {
	return r._contactSeller
}

var poolTmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterAnomalyrecourseHomedecorationResponseRequest()
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationResponseRequest 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest
func GetTmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest() *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest {
	return poolTmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest.Get().(*TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest 将 TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest(v *TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest.Put(v)
}
