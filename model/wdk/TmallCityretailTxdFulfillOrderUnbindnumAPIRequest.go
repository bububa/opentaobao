package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcityretailtxdfulfillorderunbindnumAPIRequest 淘鲜达虚拟号服务解绑接口 API请求
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
type TmallcityretailtxdfulfillorderunbindnumAPIRequest struct {
	model.Params
	// 淘鲜达交易单号
	_sourceOrderId string
	// 订阅关系id
	_subId string
}

// NewTmallcityretailtxdfulfillorderunbindnumRequest 初始化TmallcityretailtxdfulfillorderunbindnumAPIRequest对象
func NewTmallcityretailtxdfulfillorderunbindnumRequest() *TmallcityretailtxdfulfillorderunbindnumAPIRequest {
	return &TmallcityretailtxdfulfillorderunbindnumAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcityretailtxdfulfillorderunbindnumAPIRequest) GetApiMethodName() string {
	return "tmall.cityretail.txd.fulfill.order.unbindnum"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcityretailtxdfulfillorderunbindnumAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcityretailtxdfulfillorderunbindnumAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceOrderId is SourceOrderId Setter
// 淘鲜达交易单号
func (r *TmallcityretailtxdfulfillorderunbindnumAPIRequest) SetSourceOrderId(_sourceOrderId string) error {
	r._sourceOrderId = _sourceOrderId
	r.Set("source_order_id", _sourceOrderId)
	return nil
}

// GetSourceOrderId SourceOrderId Getter
func (r TmallcityretailtxdfulfillorderunbindnumAPIRequest) GetSourceOrderId() string {
	return r._sourceOrderId
}

// SetSubId is SubId Setter
// 订阅关系id
func (r *TmallcityretailtxdfulfillorderunbindnumAPIRequest) SetSubId(_subId string) error {
	r._subId = _subId
	r.Set("sub_id", _subId)
	return nil
}

// GetSubId SubId Getter
func (r TmallcityretailtxdfulfillorderunbindnumAPIRequest) GetSubId() string {
	return r._subId
}
