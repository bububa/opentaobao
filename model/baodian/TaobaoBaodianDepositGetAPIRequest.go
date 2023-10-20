package baodian

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianDepositGetAPIRequest 宝点用户帐户查询（已迁移） API请求
// taobao.baodian.deposit.get
//
// 查询用户宝点帐户信息及当前宝点价格
type TaobaoBaodianDepositGetAPIRequest struct {
	model.Params
}

// NewTaobaoBaodianDepositGetRequest 初始化TaobaoBaodianDepositGetAPIRequest对象
func NewTaobaoBaodianDepositGetRequest() *TaobaoBaodianDepositGetAPIRequest {
	return &TaobaoBaodianDepositGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaodianDepositGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaodianDepositGetAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.deposit.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaodianDepositGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaodianDepositGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoBaodianDepositGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaodianDepositGetRequest()
	},
}

// GetTaobaoBaodianDepositGetRequest 从 sync.Pool 获取 TaobaoBaodianDepositGetAPIRequest
func GetTaobaoBaodianDepositGetAPIRequest() *TaobaoBaodianDepositGetAPIRequest {
	return poolTaobaoBaodianDepositGetAPIRequest.Get().(*TaobaoBaodianDepositGetAPIRequest)
}

// ReleaseTaobaoBaodianDepositGetAPIRequest 将 TaobaoBaodianDepositGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaodianDepositGetAPIRequest(v *TaobaoBaodianDepositGetAPIRequest) {
	v.Reset()
	poolTaobaoBaodianDepositGetAPIRequest.Put(v)
}
