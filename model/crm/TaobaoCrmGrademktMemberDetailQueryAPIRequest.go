package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgrademktmemberdetailqueryAPIRequest 会员等级营销-等级营销活动查询 API请求
// taobao.crm.grademkt.member.detail.query
//
// 商家通过该接口查询等级营销活动
type TaobaocrmgrademktmemberdetailqueryAPIRequest struct {
	model.Params
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 扩展字段
	_feather string
}

// NewTaobaocrmgrademktmemberdetailqueryRequest 初始化TaobaocrmgrademktmemberdetailqueryAPIRequest对象
func NewTaobaocrmgrademktmemberdetailqueryRequest() *TaobaocrmgrademktmemberdetailqueryAPIRequest {
	return &TaobaocrmgrademktmemberdetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgrademktmemberdetailqueryAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgrademktmemberdetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgrademktmemberdetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaocrmgrademktmemberdetailqueryAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaocrmgrademktmemberdetailqueryAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 扩展字段
func (r *TaobaocrmgrademktmemberdetailqueryAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaocrmgrademktmemberdetailqueryAPIRequest) GetFeather() string {
	return r._feather
}
