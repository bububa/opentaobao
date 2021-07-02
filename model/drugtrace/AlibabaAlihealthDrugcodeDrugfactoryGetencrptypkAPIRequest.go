package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest 获取加密公钥 API请求
// alibaba.alihealth.drugcode.drugfactory.getencrptypk
//
// 获取服务端给药厂用来加密的公钥
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest struct {
	model.Params
	// 企业Id
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.getencrptypk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetRefEntId() string {
	return r._refEntId
}
