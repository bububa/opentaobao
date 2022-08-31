package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.deposit.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
