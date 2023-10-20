package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivityupdateAPIRequest 修改活动信息 API请求
// taobao.ump.activity.update
//
// 修改营销活动
type TaobaoumpactivityupdateAPIRequest struct {
	model.Params
	// 营销活动内容，json格式，通过ump sdk 的marketingTool来生成
	_content string
	// 活动id
	_actId int64
}

// NewTaobaoumpactivityupdateRequest 初始化TaobaoumpactivityupdateAPIRequest对象
func NewTaobaoumpactivityupdateRequest() *TaobaoumpactivityupdateAPIRequest {
	return &TaobaoumpactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpactivityupdateAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 营销活动内容，json格式，通过ump sdk 的marketingTool来生成
func (r *TaobaoumpactivityupdateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoumpactivityupdateAPIRequest) GetContent() string {
	return r._content
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoumpactivityupdateAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumpactivityupdateAPIRequest) GetActId() int64 {
	return r._actId
}
