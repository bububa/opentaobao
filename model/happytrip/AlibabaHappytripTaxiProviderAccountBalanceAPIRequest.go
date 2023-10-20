package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiProviderAccountBalanceAPIRequest 供应商渠道余额 API请求
// alibaba.happytrip.taxi.provider.account.balance
//
// 查询不同供应商不同渠道账户余额
type AlibabaHappytripTaxiProviderAccountBalanceAPIRequest struct {
	model.Params
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
}

// NewAlibabaHappytripTaxiProviderAccountBalanceRequest 初始化AlibabaHappytripTaxiProviderAccountBalanceAPIRequest对象
func NewAlibabaHappytripTaxiProviderAccountBalanceRequest() *AlibabaHappytripTaxiProviderAccountBalanceAPIRequest {
	return &AlibabaHappytripTaxiProviderAccountBalanceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) Reset() {
	r._costCenter = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.provider.account.balance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCostCenter is CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) SetCostCenter(_costCenter string) error {
	r._costCenter = _costCenter
	r.Set("cost_center", _costCenter)
	return nil
}

// GetCostCenter CostCenter Getter
func (r AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) GetCostCenter() string {
	return r._costCenter
}

var poolAlibabaHappytripTaxiProviderAccountBalanceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiProviderAccountBalanceRequest()
	},
}

// GetAlibabaHappytripTaxiProviderAccountBalanceRequest 从 sync.Pool 获取 AlibabaHappytripTaxiProviderAccountBalanceAPIRequest
func GetAlibabaHappytripTaxiProviderAccountBalanceAPIRequest() *AlibabaHappytripTaxiProviderAccountBalanceAPIRequest {
	return poolAlibabaHappytripTaxiProviderAccountBalanceAPIRequest.Get().(*AlibabaHappytripTaxiProviderAccountBalanceAPIRequest)
}

// ReleaseAlibabaHappytripTaxiProviderAccountBalanceAPIRequest 将 AlibabaHappytripTaxiProviderAccountBalanceAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiProviderAccountBalanceAPIRequest(v *AlibabaHappytripTaxiProviderAccountBalanceAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiProviderAccountBalanceAPIRequest.Put(v)
}
