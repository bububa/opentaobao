package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpdetailupdateAPIRequest 修改活动详情 API请求
// taobao.ump.detail.update
//
// 更新活动详情
type TaobaoumpdetailupdateAPIRequest struct {
	model.Params
	// 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容
	_content string
	// 活动详情id
	_detailId int64
}

// NewTaobaoumpdetailupdateRequest 初始化TaobaoumpdetailupdateAPIRequest对象
func NewTaobaoumpdetailupdateRequest() *TaobaoumpdetailupdateAPIRequest {
	return &TaobaoumpdetailupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpdetailupdateAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpdetailupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpdetailupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容
func (r *TaobaoumpdetailupdateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoumpdetailupdateAPIRequest) GetContent() string {
	return r._content
}

// SetDetailId is DetailId Setter
// 活动详情id
func (r *TaobaoumpdetailupdateAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoumpdetailupdateAPIRequest) GetDetailId() int64 {
	return r._detailId
}
