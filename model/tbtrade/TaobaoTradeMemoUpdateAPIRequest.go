package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradememoupdateAPIRequest 修改交易备注 API请求
// taobao.trade.memo.update
//
// 需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能
type TaobaotradememoupdateAPIRequest struct {
	model.Params
	// 卖家交易备注。最大长度: 1000个字节
	_memo string
	// 交易编号
	_tid int64
	// 卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
	// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
	_reset bool
}

// NewTaobaotradememoupdateRequest 初始化TaobaotradememoupdateAPIRequest对象
func NewTaobaotradememoupdateRequest() *TaobaotradememoupdateAPIRequest {
	return &TaobaotradememoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradememoupdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.memo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradememoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradememoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 卖家交易备注。最大长度: 1000个字节
func (r *TaobaotradememoupdateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaotradememoupdateAPIRequest) GetMemo() string {
	return r._memo
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaotradememoupdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradememoupdateAPIRequest) GetTid() int64 {
	return r._tid
}

// SetFlag is Flag Setter
// 卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaotradememoupdateAPIRequest) SetFlag(_flag int64) error {
	r._flag = _flag
	r.Set("flag", _flag)
	return nil
}

// GetFlag Flag Getter
func (r TaobaotradememoupdateAPIRequest) GetFlag() int64 {
	return r._flag
}

// SetReset is Reset Setter
// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
func (r *TaobaotradememoupdateAPIRequest) SetReset(_reset bool) error {
	r._reset = _reset
	r.Set("reset", _reset)
	return nil
}

// GetReset Reset Getter
func (r TaobaotradememoupdateAPIRequest) GetReset() bool {
	return r._reset
}
