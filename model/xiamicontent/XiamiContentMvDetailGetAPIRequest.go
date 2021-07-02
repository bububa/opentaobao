package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentMvDetailGetAPIRequest 获取mv详情 API请求
// xiami.content.mv.detail.get
//
// 获取mv详情
type XiamiContentMvDetailGetAPIRequest struct {
	model.Params
	// mvId
	_mvIds []int64
}

// NewXiamiContentMvDetailGetRequest 初始化XiamiContentMvDetailGetAPIRequest对象
func NewXiamiContentMvDetailGetRequest() *XiamiContentMvDetailGetAPIRequest {
	return &XiamiContentMvDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentMvDetailGetAPIRequest) GetApiMethodName() string {
	return "xiami.content.mv.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentMvDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MvIds Setter
// mvId
func (r *XiamiContentMvDetailGetAPIRequest) SetMvIds(_mvIds []int64) error {
	r._mvIds = _mvIds
	r.Set("mv_ids", _mvIds)
	return nil
}

// Get MvIds Getter
func (r XiamiContentMvDetailGetAPIRequest) GetMvIds() []int64 {
	return r._mvIds
}
