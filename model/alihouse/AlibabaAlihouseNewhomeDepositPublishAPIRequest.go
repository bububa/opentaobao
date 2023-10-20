package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeDepositPublishAPIRequest 创建、修改、发布房产预存金商品 API请求
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
type AlibabaAlihouseNewhomeDepositPublishAPIRequest struct {
	model.Params
	// 房产预存金属性
	_preDepositGoldSaveParam *PreDepositGoldSaveParam
}

// NewAlibabaAlihouseNewhomeDepositPublishRequest 初始化AlibabaAlihouseNewhomeDepositPublishAPIRequest对象
func NewAlibabaAlihouseNewhomeDepositPublishRequest() *AlibabaAlihouseNewhomeDepositPublishAPIRequest {
	return &AlibabaAlihouseNewhomeDepositPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeDepositPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.deposit.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeDepositPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeDepositPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreDepositGoldSaveParam is PreDepositGoldSaveParam Setter
// 房产预存金属性
func (r *AlibabaAlihouseNewhomeDepositPublishAPIRequest) SetPreDepositGoldSaveParam(_preDepositGoldSaveParam *PreDepositGoldSaveParam) error {
	r._preDepositGoldSaveParam = _preDepositGoldSaveParam
	r.Set("pre_deposit_gold_save_param", _preDepositGoldSaveParam)
	return nil
}

// GetPreDepositGoldSaveParam PreDepositGoldSaveParam Getter
func (r AlibabaAlihouseNewhomeDepositPublishAPIRequest) GetPreDepositGoldSaveParam() *PreDepositGoldSaveParam {
	return r._preDepositGoldSaveParam
}
