package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityUpdateAPIRequest 修改活动信息 API请求
// taobao.ump.activity.update
//
// 修改营销活动
type TaobaoUmpActivityUpdateAPIRequest struct {
	model.Params
	// 营销活动内容，json格式，通过ump sdk 的marketingTool来生成
	_content string
	// 活动id
	_actId int64
}

// NewTaobaoUmpActivityUpdateRequest 初始化TaobaoUmpActivityUpdateAPIRequest对象
func NewTaobaoUmpActivityUpdateRequest() *TaobaoUmpActivityUpdateAPIRequest {
	return &TaobaoUmpActivityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetContent is Content Setter
// 营销活动内容，json格式，通过ump sdk 的marketingTool来生成
func (r *TaobaoUmpActivityUpdateAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoUmpActivityUpdateAPIRequest) GetContent() string {
	return r._content
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpActivityUpdateAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpActivityUpdateAPIRequest) GetActId() int64 {
	return r._actId
}
