package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotraderateexplainaddAPIRequest 商城评价解释接口 API请求
// taobao.traderate.explain.add
//
// 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
type TaobaotraderateexplainaddAPIRequest struct {
	model.Params
	// 评价解释内容，最大长度：500个汉字
	_reply string
	// 子订单ID
	_oid int64
}

// NewTaobaotraderateexplainaddRequest 初始化TaobaotraderateexplainaddAPIRequest对象
func NewTaobaotraderateexplainaddRequest() *TaobaotraderateexplainaddAPIRequest {
	return &TaobaotraderateexplainaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotraderateexplainaddAPIRequest) GetApiMethodName() string {
	return "taobao.traderate.explain.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotraderateexplainaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotraderateexplainaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReply is Reply Setter
// 评价解释内容，最大长度：500个汉字
func (r *TaobaotraderateexplainaddAPIRequest) SetReply(_reply string) error {
	r._reply = _reply
	r.Set("reply", _reply)
	return nil
}

// GetReply Reply Getter
func (r TaobaotraderateexplainaddAPIRequest) GetReply() string {
	return r._reply
}

// SetOid is Oid Setter
// 子订单ID
func (r *TaobaotraderateexplainaddAPIRequest) SetOid(_oid int64) error {
	r._oid = _oid
	r.Set("oid", _oid)
	return nil
}

// GetOid Oid Getter
func (r TaobaotraderateexplainaddAPIRequest) GetOid() int64 {
	return r._oid
}
