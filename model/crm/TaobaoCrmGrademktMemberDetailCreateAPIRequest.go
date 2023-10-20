package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgrademktmemberdetailcreateAPIRequest 会员等级营销-创建商品等级营销明细 API请求
// taobao.crm.grademkt.member.detail.create
//
// 创建商品等级营销明细
type TaobaocrmgrademktmemberdetailcreateAPIRequest struct {
	model.Params
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 扩展字段
	_feather string
}

// NewTaobaocrmgrademktmemberdetailcreateRequest 初始化TaobaocrmgrademktmemberdetailcreateAPIRequest对象
func NewTaobaocrmgrademktmemberdetailcreateRequest() *TaobaocrmgrademktmemberdetailcreateAPIRequest {
	return &TaobaocrmgrademktmemberdetailcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgrademktmemberdetailcreateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgrademktmemberdetailcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgrademktmemberdetailcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaocrmgrademktmemberdetailcreateAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaocrmgrademktmemberdetailcreateAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 扩展字段
func (r *TaobaocrmgrademktmemberdetailcreateAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaocrmgrademktmemberdetailcreateAPIRequest) GetFeather() string {
	return r._feather
}
