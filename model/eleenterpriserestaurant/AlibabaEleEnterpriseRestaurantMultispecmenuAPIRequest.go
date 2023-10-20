package eleenterpriserestaurant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest 查询餐厅菜单 API请求
// alibaba.ele.enterprise.restaurant.multispecmenu
//
// 查询餐厅菜单
type AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest struct {
	model.Params
	// 餐厅ID
	_erestaurantId string
}

// NewAlibabaEleEnterpriseRestaurantMultispecmenuRequest 初始化AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest对象
func NewAlibabaEleEnterpriseRestaurantMultispecmenuRequest() *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest {
	return &AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) Reset() {
	r._erestaurantId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.restaurant.multispecmenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErestaurantId is ErestaurantId Setter
// 餐厅ID
func (r *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) SetErestaurantId(_erestaurantId string) error {
	r._erestaurantId = _erestaurantId
	r.Set("erestaurant_id", _erestaurantId)
	return nil
}

// GetErestaurantId ErestaurantId Getter
func (r AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) GetErestaurantId() string {
	return r._erestaurantId
}

var poolAlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseRestaurantMultispecmenuRequest()
	},
}

// GetAlibabaEleEnterpriseRestaurantMultispecmenuRequest 从 sync.Pool 获取 AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest
func GetAlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest() *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest {
	return poolAlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest.Get().(*AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest)
}

// ReleaseAlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest 将 AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest(v *AlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseRestaurantMultispecmenuAPIRequest.Put(v)
}
