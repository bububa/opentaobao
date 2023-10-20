package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotraderatelistaddAPIRequest 针对父子订单新增批量评价 API请求
// taobao.traderate.list.add
//
// 针对父子订单新增批量评价(&lt;font color=&#34;red&#34;&gt;注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价&lt;/font&gt;)
type TaobaotraderatelistaddAPIRequest struct {
	model.Params
	// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
	_result string
	// 评价者角色。可选值:seller(卖家),buyer(买家)
	_role string
	// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
	_content string
	// 交易ID
	_tid int64
	// 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名
	_anony bool
}

// NewTaobaotraderatelistaddRequest 初始化TaobaotraderatelistaddAPIRequest对象
func NewTaobaotraderatelistaddRequest() *TaobaotraderatelistaddAPIRequest {
	return &TaobaotraderatelistaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotraderatelistaddAPIRequest) GetApiMethodName() string {
	return "taobao.traderate.list.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotraderatelistaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotraderatelistaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResult is Result Setter
// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaotraderatelistaddAPIRequest) SetResult(_result string) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r TaobaotraderatelistaddAPIRequest) GetResult() string {
	return r._result
}

// SetRole is Role Setter
// 评价者角色。可选值:seller(卖家),buyer(买家)
func (r *TaobaotraderatelistaddAPIRequest) SetRole(_role string) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r TaobaotraderatelistaddAPIRequest) GetRole() string {
	return r._role
}

// SetContent is Content Setter
// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
func (r *TaobaotraderatelistaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaotraderatelistaddAPIRequest) GetContent() string {
	return r._content
}

// SetTid is Tid Setter
// 交易ID
func (r *TaobaotraderatelistaddAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotraderatelistaddAPIRequest) GetTid() int64 {
	return r._tid
}

// SetAnony is Anony Setter
// 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名
func (r *TaobaotraderatelistaddAPIRequest) SetAnony(_anony bool) error {
	r._anony = _anony
	r.Set("anony", _anony)
	return nil
}

// GetAnony Anony Getter
func (r TaobaotraderatelistaddAPIRequest) GetAnony() bool {
	return r._anony
}
