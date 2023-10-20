package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptraveltradememoupdateAPIRequest 修改一笔交易备注 API请求
// taobao.alitrip.travel.trade.memo.update
//
// 更新一笔交易备注
type TaobaoalitriptraveltradememoupdateAPIRequest struct {
	model.Params
	// 交易备注。最大长度: 1000个字节
	_memo string
	// 交易ID
	_tid int64
	// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
	// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
	_reset bool
}

// NewTaobaoalitriptraveltradememoupdateRequest 初始化TaobaoalitriptraveltradememoupdateAPIRequest对象
func NewTaobaoalitriptraveltradememoupdateRequest() *TaobaoalitriptraveltradememoupdateAPIRequest {
	return &TaobaoalitriptraveltradememoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.trade.memo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaoalitriptraveltradememoupdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetMemo() string {
	return r._memo
}

// SetTid is Tid Setter
// 交易ID
func (r *TaobaoalitriptraveltradememoupdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetTid() int64 {
	return r._tid
}

// SetFlag is Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaoalitriptraveltradememoupdateAPIRequest) SetFlag(_flag int64) error {
	r._flag = _flag
	r.Set("flag", _flag)
	return nil
}

// GetFlag Flag Getter
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetFlag() int64 {
	return r._flag
}

// SetReset is Reset Setter
// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
func (r *TaobaoalitriptraveltradememoupdateAPIRequest) SetReset(_reset bool) error {
	r._reset = _reset
	r.Set("reset", _reset)
	return nil
}

// GetReset Reset Getter
func (r TaobaoalitriptraveltradememoupdateAPIRequest) GetReset() bool {
	return r._reset
}
