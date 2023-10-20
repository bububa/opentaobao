package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgrademktmemberdetaildeleteAPIRequest 会员等级营销-删除商品等级营销明细 API请求
// taobao.crm.grademkt.member.detail.delete
//
// 删除商品等级营销明细
type TaobaocrmgrademktmemberdetaildeleteAPIRequest struct {
	model.Params
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 扩展字段
	_feather string
}

// NewTaobaocrmgrademktmemberdetaildeleteRequest 初始化TaobaocrmgrademktmemberdetaildeleteAPIRequest对象
func NewTaobaocrmgrademktmemberdetaildeleteRequest() *TaobaocrmgrademktmemberdetaildeleteAPIRequest {
	return &TaobaocrmgrademktmemberdetaildeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgrademktmemberdetaildeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgrademktmemberdetaildeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgrademktmemberdetaildeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaocrmgrademktmemberdetaildeleteAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaocrmgrademktmemberdetaildeleteAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 扩展字段
func (r *TaobaocrmgrademktmemberdetaildeleteAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaocrmgrademktmemberdetaildeleteAPIRequest) GetFeather() string {
	return r._feather
}
