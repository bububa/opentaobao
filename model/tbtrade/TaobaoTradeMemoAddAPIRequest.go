package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradememoaddAPIRequest 对一笔交易添加备注 API请求
// taobao.trade.memo.add
//
// 根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update
type TaobaotradememoaddAPIRequest struct {
	model.Params
	// 交易备注。最大长度: 1000个字节
	_memo string
	// 交易编号
	_tid int64
	// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
}

// NewTaobaotradememoaddRequest 初始化TaobaotradememoaddAPIRequest对象
func NewTaobaotradememoaddRequest() *TaobaotradememoaddAPIRequest {
	return &TaobaotradememoaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradememoaddAPIRequest) GetApiMethodName() string {
	return "taobao.trade.memo.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradememoaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradememoaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaotradememoaddAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaotradememoaddAPIRequest) GetMemo() string {
	return r._memo
}

// SetTid is Tid Setter
// 交易编号
func (r *TaobaotradememoaddAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradememoaddAPIRequest) GetTid() int64 {
	return r._tid
}

// SetFlag is Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaotradememoaddAPIRequest) SetFlag(_flag int64) error {
	r._flag = _flag
	r.Set("flag", _flag)
	return nil
}

// GetFlag Flag Getter
func (r TaobaotradememoaddAPIRequest) GetFlag() int64 {
	return r._flag
}
