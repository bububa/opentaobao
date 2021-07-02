package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintSingleCustomareaGetAPIRequest 获取商家单一自定义区 API请求
// cainiao.cloudprint.single.customarea.get
//
// 商家所有快递公司模板只有一个自定义区
type CainiaoCloudprintSingleCustomareaGetAPIRequest struct {
	model.Params
	// 这是商家用户id
	_sellerId int64
}

// NewCainiaoCloudprintSingleCustomareaGetRequest 初始化CainiaoCloudprintSingleCustomareaGetAPIRequest对象
func NewCainiaoCloudprintSingleCustomareaGetRequest() *CainiaoCloudprintSingleCustomareaGetAPIRequest {
	return &CainiaoCloudprintSingleCustomareaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintSingleCustomareaGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.single.customarea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintSingleCustomareaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSellerId is SellerId Setter
// 这是商家用户id
func (r *CainiaoCloudprintSingleCustomareaGetAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r CainiaoCloudprintSingleCustomareaGetAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
