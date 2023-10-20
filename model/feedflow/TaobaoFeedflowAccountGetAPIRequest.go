package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowAccountGetAPIRequest 获取信息流账户详情 API请求
// taobao.feedflow.account.get
//
// 获取账户信息接口。
// (1) BP显示余额 (字段 ：banlance ) = 现金余额(字段：cash_balance) + 赠款余额；
// (2) 可用余额(字段：availableBalance) = BP显示余额
// (3) 红包(字段：redPacket)
type TaobaoFeedflowAccountGetAPIRequest struct {
	model.Params
}

// NewTaobaoFeedflowAccountGetRequest 初始化TaobaoFeedflowAccountGetAPIRequest对象
func NewTaobaoFeedflowAccountGetRequest() *TaobaoFeedflowAccountGetAPIRequest {
	return &TaobaoFeedflowAccountGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowAccountGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountGetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowAccountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoFeedflowAccountGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowAccountGetRequest()
	},
}

// GetTaobaoFeedflowAccountGetRequest 从 sync.Pool 获取 TaobaoFeedflowAccountGetAPIRequest
func GetTaobaoFeedflowAccountGetAPIRequest() *TaobaoFeedflowAccountGetAPIRequest {
	return poolTaobaoFeedflowAccountGetAPIRequest.Get().(*TaobaoFeedflowAccountGetAPIRequest)
}

// ReleaseTaobaoFeedflowAccountGetAPIRequest 将 TaobaoFeedflowAccountGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowAccountGetAPIRequest(v *TaobaoFeedflowAccountGetAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowAccountGetAPIRequest.Put(v)
}
