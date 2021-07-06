package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaChongzhiQueryecardsAPIRequest 查询指定商家的可用的话费宝贝接口 API请求
// alibaba.chongzhi.queryecards
//
// 查询指定商家的可用的话费宝贝
type AlibabaChongzhiQueryecardsAPIRequest struct {
	model.Params
	// 来源
	_clientSource string
	// 号码
	_mobile int64
}

// NewAlibabaChongzhiQueryecardsRequest 初始化AlibabaChongzhiQueryecardsAPIRequest对象
func NewAlibabaChongzhiQueryecardsRequest() *AlibabaChongzhiQueryecardsAPIRequest {
	return &AlibabaChongzhiQueryecardsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaChongzhiQueryecardsAPIRequest) GetApiMethodName() string {
	return "alibaba.chongzhi.queryecards"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaChongzhiQueryecardsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetClientSource is ClientSource Setter
// 来源
func (r *AlibabaChongzhiQueryecardsAPIRequest) SetClientSource(_clientSource string) error {
	r._clientSource = _clientSource
	r.Set("client_source", _clientSource)
	return nil
}

// GetClientSource ClientSource Getter
func (r AlibabaChongzhiQueryecardsAPIRequest) GetClientSource() string {
	return r._clientSource
}

// SetMobile is Mobile Setter
// 号码
func (r *AlibabaChongzhiQueryecardsAPIRequest) SetMobile(_mobile int64) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaChongzhiQueryecardsAPIRequest) GetMobile() int64 {
	return r._mobile
}
