package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradepostageupdateAPIRequest 修改交易邮费价格 API请求
// taobao.trade.postage.update
//
// 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
// &lt;br/&gt; &lt;span style=&#34;color:red&#34;&gt; API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750&lt;/span&gt;
type TaobaotradepostageupdateAPIRequest struct {
	model.Params
	// 邮费价格(邮费单位是元）
	_postFee string
	// 主订单编号
	_tid int64
}

// NewTaobaotradepostageupdateRequest 初始化TaobaotradepostageupdateAPIRequest对象
func NewTaobaotradepostageupdateRequest() *TaobaotradepostageupdateAPIRequest {
	return &TaobaotradepostageupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradepostageupdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.postage.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradepostageupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradepostageupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPostFee is PostFee Setter
// 邮费价格(邮费单位是元）
func (r *TaobaotradepostageupdateAPIRequest) SetPostFee(_postFee string) error {
	r._postFee = _postFee
	r.Set("post_fee", _postFee)
	return nil
}

// GetPostFee PostFee Getter
func (r TaobaotradepostageupdateAPIRequest) GetPostFee() string {
	return r._postFee
}

// SetTid is Tid Setter
// 主订单编号
func (r *TaobaotradepostageupdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradepostageupdateAPIRequest) GetTid() int64 {
	return r._tid
}
