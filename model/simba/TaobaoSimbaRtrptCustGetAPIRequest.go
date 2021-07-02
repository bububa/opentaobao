package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptCustGetAPIRequest 获取账户实时报表数据 API请求
// taobao.simba.rtrpt.cust.get
//
// 获取账户实时报表数据
type TaobaoSimbaRtrptCustGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// NewTaobaoSimbaRtrptCustGetRequest 初始化TaobaoSimbaRtrptCustGetAPIRequest对象
func NewTaobaoSimbaRtrptCustGetRequest() *TaobaoSimbaRtrptCustGetAPIRequest {
	return &TaobaoSimbaRtrptCustGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCustGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.rtrpt.cust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCustGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaRtrptCustGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaRtrptCustGetAPIRequest) GetNick() string {
	return r._nick
}

// SetTheDate is TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCustGetAPIRequest) SetTheDate(_theDate string) error {
	r._theDate = _theDate
	r.Set("the_date", _theDate)
	return nil
}

// GetTheDate TheDate Getter
func (r TaobaoSimbaRtrptCustGetAPIRequest) GetTheDate() string {
	return r._theDate
}
