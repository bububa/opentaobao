package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) Reset() {
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.getencrptypk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业Id
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest()
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest
func GetAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest {
	return poolAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest.Get().(*AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest 将 AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest(v *AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest.Put(v)
}
