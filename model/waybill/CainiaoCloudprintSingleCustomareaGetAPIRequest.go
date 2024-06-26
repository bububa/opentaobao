package waybill

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintSingleCustomareaGetAPIRequest) Reset() {
	r._sellerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintSingleCustomareaGetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.single.customarea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintSingleCustomareaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintSingleCustomareaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoCloudprintSingleCustomareaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintSingleCustomareaGetRequest()
	},
}

// GetCainiaoCloudprintSingleCustomareaGetRequest 从 sync.Pool 获取 CainiaoCloudprintSingleCustomareaGetAPIRequest
func GetCainiaoCloudprintSingleCustomareaGetAPIRequest() *CainiaoCloudprintSingleCustomareaGetAPIRequest {
	return poolCainiaoCloudprintSingleCustomareaGetAPIRequest.Get().(*CainiaoCloudprintSingleCustomareaGetAPIRequest)
}

// ReleaseCainiaoCloudprintSingleCustomareaGetAPIRequest 将 CainiaoCloudprintSingleCustomareaGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintSingleCustomareaGetAPIRequest(v *CainiaoCloudprintSingleCustomareaGetAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintSingleCustomareaGetAPIRequest.Put(v)
}
