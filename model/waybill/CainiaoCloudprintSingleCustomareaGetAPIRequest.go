package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintsinglecustomareagetAPIRequest 获取商家单一自定义区 API请求
// cainiao.cloudprint.single.customarea.get
//
// 商家所有快递公司模板只有一个自定义区
type CainiaocloudprintsinglecustomareagetAPIRequest struct {
	model.Params
	// 这是商家用户id
	_sellerId int64
}

// NewCainiaocloudprintsinglecustomareagetRequest 初始化CainiaocloudprintsinglecustomareagetAPIRequest对象
func NewCainiaocloudprintsinglecustomareagetRequest() *CainiaocloudprintsinglecustomareagetAPIRequest {
	return &CainiaocloudprintsinglecustomareagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintsinglecustomareagetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.single.customarea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintsinglecustomareagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintsinglecustomareagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerId is SellerId Setter
// 这是商家用户id
func (r *CainiaocloudprintsinglecustomareagetAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r CainiaocloudprintsinglecustomareagetAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
