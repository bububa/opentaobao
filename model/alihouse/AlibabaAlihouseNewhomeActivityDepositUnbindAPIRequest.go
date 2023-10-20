package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest 销售活动解绑预存金商品 API请求
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
type AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest struct {
	model.Params
	// 销售活动解绑预存金报文
	_preDepositGoldUnbindDto *PreDepositGoldUnbindDto
}

// NewAlibabaAlihouseNewhomeActivityDepositUnbindRequest 初始化AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest对象
func NewAlibabaAlihouseNewhomeActivityDepositUnbindRequest() *AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest {
	return &AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) Reset() {
	r._preDepositGoldUnbindDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.deposit.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreDepositGoldUnbindDto is PreDepositGoldUnbindDto Setter
// 销售活动解绑预存金报文
func (r *AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) SetPreDepositGoldUnbindDto(_preDepositGoldUnbindDto *PreDepositGoldUnbindDto) error {
	r._preDepositGoldUnbindDto = _preDepositGoldUnbindDto
	r.Set("pre_deposit_gold_unbind_dto", _preDepositGoldUnbindDto)
	return nil
}

// GetPreDepositGoldUnbindDto PreDepositGoldUnbindDto Getter
func (r AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) GetPreDepositGoldUnbindDto() *PreDepositGoldUnbindDto {
	return r._preDepositGoldUnbindDto
}

var poolAlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeActivityDepositUnbindRequest()
	},
}

// GetAlibabaAlihouseNewhomeActivityDepositUnbindRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest
func GetAlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest() *AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest {
	return poolAlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest.Get().(*AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest 将 AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest(v *AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest.Put(v)
}
