package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpActivityUpdateAPIRequest) Reset() {
	r._content = ""
	r._actId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoUmpActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpActivityUpdateRequest()
	},
}

// GetTaobaoUmpActivityUpdateRequest 从 sync.Pool 获取 TaobaoUmpActivityUpdateAPIRequest
func GetTaobaoUmpActivityUpdateAPIRequest() *TaobaoUmpActivityUpdateAPIRequest {
	return poolTaobaoUmpActivityUpdateAPIRequest.Get().(*TaobaoUmpActivityUpdateAPIRequest)
}

// ReleaseTaobaoUmpActivityUpdateAPIRequest 将 TaobaoUmpActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpActivityUpdateAPIRequest(v *TaobaoUmpActivityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoUmpActivityUpdateAPIRequest.Put(v)
}
