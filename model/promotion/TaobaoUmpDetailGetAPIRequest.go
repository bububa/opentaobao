package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpdetailgetAPIRequest 查询活动详情 API请求
// taobao.ump.detail.get
//
// 查询活动详情
type TaobaoumpdetailgetAPIRequest struct {
	model.Params
	// 活动详情的id
	_detailId int64
}

// NewTaobaoumpdetailgetRequest 初始化TaobaoumpdetailgetAPIRequest对象
func NewTaobaoumpdetailgetRequest() *TaobaoumpdetailgetAPIRequest {
	return &TaobaoumpdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpdetailgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailId is DetailId Setter
// 活动详情的id
func (r *TaobaoumpdetailgetAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoumpdetailgetAPIRequest) GetDetailId() int64 {
	return r._detailId
}
