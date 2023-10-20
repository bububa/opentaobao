package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomedepositpublishAPIRequest 创建、修改、发布房产预存金商品 API请求
// alibaba.alihouse.newhome.deposit.publish
//
// 创建、修改、发布房产预存金商品
type AlibabaalihousenewhomedepositpublishAPIRequest struct {
	model.Params
	// 房产预存金属性
	_preDepositGoldSaveParam *PreDepositGoldSaveParam
}

// NewAlibabaalihousenewhomedepositpublishRequest 初始化AlibabaalihousenewhomedepositpublishAPIRequest对象
func NewAlibabaalihousenewhomedepositpublishRequest() *AlibabaalihousenewhomedepositpublishAPIRequest {
	return &AlibabaalihousenewhomedepositpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomedepositpublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.deposit.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomedepositpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomedepositpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreDepositGoldSaveParam is PreDepositGoldSaveParam Setter
// 房产预存金属性
func (r *AlibabaalihousenewhomedepositpublishAPIRequest) SetPreDepositGoldSaveParam(_preDepositGoldSaveParam *PreDepositGoldSaveParam) error {
	r._preDepositGoldSaveParam = _preDepositGoldSaveParam
	r.Set("pre_deposit_gold_save_param", _preDepositGoldSaveParam)
	return nil
}

// GetPreDepositGoldSaveParam PreDepositGoldSaveParam Getter
func (r AlibabaalihousenewhomedepositpublishAPIRequest) GetPreDepositGoldSaveParam() *PreDepositGoldSaveParam {
	return r._preDepositGoldSaveParam
}
