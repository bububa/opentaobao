package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotraderateaddAPIRequest 新增单个评价 API请求
// taobao.traderate.add
//
// 新增单个评价(&lt;font color=&#34;red&#34;&gt;注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价&lt;/font&gt;)
type TaobaotraderateaddAPIRequest struct {
	model.Params
	// 评价结果,可选值:good(好评),neutral(中评),bad(差评)
	_result string
	// 评价者角色,可选值:seller(卖家),buyer(买家)
	_role string
	// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
	_content string
	// 交易ID
	_tid int64
	// 子订单ID
	_oid int64
	// 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名
	_anony bool
}

// NewTaobaotraderateaddRequest 初始化TaobaotraderateaddAPIRequest对象
func NewTaobaotraderateaddRequest() *TaobaotraderateaddAPIRequest {
	return &TaobaotraderateaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotraderateaddAPIRequest) GetApiMethodName() string {
	return "taobao.traderate.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotraderateaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotraderateaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResult is Result Setter
// 评价结果,可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaotraderateaddAPIRequest) SetResult(_result string) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r TaobaotraderateaddAPIRequest) GetResult() string {
	return r._result
}

// SetRole is Role Setter
// 评价者角色,可选值:seller(卖家),buyer(买家)
func (r *TaobaotraderateaddAPIRequest) SetRole(_role string) error {
	r._role = _role
	r.Set("role", _role)
	return nil
}

// GetRole Role Getter
func (r TaobaotraderateaddAPIRequest) GetRole() string {
	return r._role
}

// SetContent is Content Setter
// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
func (r *TaobaotraderateaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaotraderateaddAPIRequest) GetContent() string {
	return r._content
}

// SetTid is Tid Setter
// 交易ID
func (r *TaobaotraderateaddAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotraderateaddAPIRequest) GetTid() int64 {
	return r._tid
}

// SetOid is Oid Setter
// 子订单ID
func (r *TaobaotraderateaddAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaotraderateaddAPIRequest) GetOid() int64 {
	return r._oid
}

// SetAnony is Anony Setter
// 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名
func (r *TaobaotraderateaddAPIRequest) SetAnony(_anony bool) error {
	r._anony = _anony
	r.Set("anony", _anony)
	return nil
}

// GetAnony Anony Getter
func (r TaobaotraderateaddAPIRequest) GetAnony() bool {
	return r._anony
}
