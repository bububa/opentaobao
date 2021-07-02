package feedflow

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountGetAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
