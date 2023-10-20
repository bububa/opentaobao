package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest 天猫服务平台商家投诉单服务商申诉接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
type TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest struct {
	model.Params
	// 申诉内容
	_remark string
	// 图片凭证
	_images string
	// 投诉单id
	_id int64
}

// NewTmallServicecenterAnomalyrecourseHomedecorationAppealRequest 初始化TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest对象
func NewTmallServicecenterAnomalyrecourseHomedecorationAppealRequest() *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest {
	return &TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) Reset() {
	r._remark = ""
	r._images = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.appeal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 申诉内容
func (r *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) GetRemark() string {
	return r._remark
}

// SetImages is Images Setter
// 图片凭证
func (r *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) SetImages(_images string) error {
	r._images = _images
	r.Set("images", _images)
	return nil
}

// GetImages Images Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) GetImages() string {
	return r._images
}

// SetId is Id Setter
// 投诉单id
func (r *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) GetId() int64 {
	return r._id
}

var poolTmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterAnomalyrecourseHomedecorationAppealRequest()
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationAppealRequest 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest
func GetTmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest() *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest {
	return poolTmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest.Get().(*TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest 将 TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest(v *TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest.Put(v)
}
