package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailGetAPIRequest 查询活动详情 API请求
// taobao.ump.detail.get
//
// 查询活动详情
type TaobaoUmpDetailGetAPIRequest struct {
	model.Params
	// 活动详情的id
	_detailId int64
}

// NewTaobaoUmpDetailGetRequest 初始化TaobaoUmpDetailGetAPIRequest对象
func NewTaobaoUmpDetailGetRequest() *TaobaoUmpDetailGetAPIRequest {
	return &TaobaoUmpDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDetailId is DetailId Setter
// 活动详情的id
func (r *TaobaoUmpDetailGetAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoUmpDetailGetAPIRequest) GetDetailId() int64 {
	return r._detailId
}
