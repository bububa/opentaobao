package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivityDepositBindAPIRequest 销售活动绑定预存金商品id API请求
// alibaba.alihouse.newhome.activity.deposit.bind
//
// 销售活动绑定预存金商品id
type AlibabaAlihouseNewhomeActivityDepositBindAPIRequest struct {
	model.Params
	// 绑定预存金商品属性
	_preDepositGoldDto *PreDepositGoldDto
}

// NewAlibabaAlihouseNewhomeActivityDepositBindRequest 初始化AlibabaAlihouseNewhomeActivityDepositBindAPIRequest对象
func NewAlibabaAlihouseNewhomeActivityDepositBindRequest() *AlibabaAlihouseNewhomeActivityDepositBindAPIRequest {
	return &AlibabaAlihouseNewhomeActivityDepositBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeActivityDepositBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.activity.deposit.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeActivityDepositBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeActivityDepositBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreDepositGoldDto is PreDepositGoldDto Setter
// 绑定预存金商品属性
func (r *AlibabaAlihouseNewhomeActivityDepositBindAPIRequest) SetPreDepositGoldDto(_preDepositGoldDto *PreDepositGoldDto) error {
	r._preDepositGoldDto = _preDepositGoldDto
	r.Set("pre_deposit_gold_dto", _preDepositGoldDto)
	return nil
}

// GetPreDepositGoldDto PreDepositGoldDto Getter
func (r AlibabaAlihouseNewhomeActivityDepositBindAPIRequest) GetPreDepositGoldDto() *PreDepositGoldDto {
	return r._preDepositGoldDto
}
