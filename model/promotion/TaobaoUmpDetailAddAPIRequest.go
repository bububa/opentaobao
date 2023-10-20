package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailAddAPIRequest 新增活动详情 API请求
// taobao.ump.detail.add
//
// 增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。
type TaobaoUmpDetailAddAPIRequest struct {
	model.Params
	// 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理
	_content string
	// 增加工具详情
	_actId int64
}

// NewTaobaoUmpDetailAddRequest 初始化TaobaoUmpDetailAddAPIRequest对象
func NewTaobaoUmpDetailAddRequest() *TaobaoUmpDetailAddAPIRequest {
	return &TaobaoUmpDetailAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpDetailAddAPIRequest) Reset() {
	r._content = ""
	r._actId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailAddAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpDetailAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理
func (r *TaobaoUmpDetailAddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoUmpDetailAddAPIRequest) GetContent() string {
	return r._content
}

// SetActId is ActId Setter
// 增加工具详情
func (r *TaobaoUmpDetailAddAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpDetailAddAPIRequest) GetActId() int64 {
	return r._actId
}

var poolTaobaoUmpDetailAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpDetailAddRequest()
	},
}

// GetTaobaoUmpDetailAddRequest 从 sync.Pool 获取 TaobaoUmpDetailAddAPIRequest
func GetTaobaoUmpDetailAddAPIRequest() *TaobaoUmpDetailAddAPIRequest {
	return poolTaobaoUmpDetailAddAPIRequest.Get().(*TaobaoUmpDetailAddAPIRequest)
}

// ReleaseTaobaoUmpDetailAddAPIRequest 将 TaobaoUmpDetailAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpDetailAddAPIRequest(v *TaobaoUmpDetailAddAPIRequest) {
	v.Reset()
	poolTaobaoUmpDetailAddAPIRequest.Put(v)
}
