package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptraveltradememoaddAPIRequest 添加一笔交易备注 API请求
// taobao.alitrip.travel.trade.memo.add
//
// 对一笔交易添加备注
type TaobaoalitriptraveltradememoaddAPIRequest struct {
	model.Params
	// 交易备注。最大长度: 1000个字节
	_memo string
	// 交易ID
	_tid int64
	// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
}

// NewTaobaoalitriptraveltradememoaddRequest 初始化TaobaoalitriptraveltradememoaddAPIRequest对象
func NewTaobaoalitriptraveltradememoaddRequest() *TaobaoalitriptraveltradememoaddAPIRequest {
	return &TaobaoalitriptraveltradememoaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptraveltradememoaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.trade.memo.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptraveltradememoaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptraveltradememoaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 交易备注。最大长度: 1000个字节
func (r *TaobaoalitriptraveltradememoaddAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoalitriptraveltradememoaddAPIRequest) GetMemo() string {
	return r._memo
}

// SetTid is Tid Setter
// 交易ID
func (r *TaobaoalitriptraveltradememoaddAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoalitriptraveltradememoaddAPIRequest) GetTid() int64 {
	return r._tid
}

// SetFlag is Flag Setter
// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
func (r *TaobaoalitriptraveltradememoaddAPIRequest) SetFlag(_flag int64) error {
	r._flag = _flag
	r.Set("flag", _flag)
	return nil
}

// GetFlag Flag Getter
func (r TaobaoalitriptraveltradememoaddAPIRequest) GetFlag() int64 {
	return r._flag
}
