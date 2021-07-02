package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTraderateExplainAddAPIRequest 商城评价解释接口 API请求
// taobao.traderate.explain.add
//
// 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
type TaobaoTraderateExplainAddAPIRequest struct {
	model.Params
	// 子订单ID
	_oid int64
	// 评价解释内容，最大长度：500个汉字
	_reply string
}

// NewTaobaoTraderateExplainAddRequest 初始化TaobaoTraderateExplainAddAPIRequest对象
func NewTaobaoTraderateExplainAddRequest() *TaobaoTraderateExplainAddAPIRequest {
	return &TaobaoTraderateExplainAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTraderateExplainAddAPIRequest) GetApiMethodName() string {
	return "taobao.traderate.explain.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTraderateExplainAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Oid Setter
// 子订单ID
func (r *TaobaoTraderateExplainAddAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// Get Oid Getter
func (r TaobaoTraderateExplainAddAPIRequest) GetOid() int64 {
	return r._oid
}

// Set is Reply Setter
// 评价解释内容，最大长度：500个汉字
func (r *TaobaoTraderateExplainAddAPIRequest) SetReply(_reply string) error {
	r._reply = _reply
	r.Set("reply", _reply)
	return nil
}

// Get Reply Getter
func (r TaobaoTraderateExplainAddAPIRequest) GetReply() string {
	return r._reply
}
