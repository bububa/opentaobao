package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest 导出所有项目的药物属性和药品信息 API请求
// alibaba.alihealth.drugcode.drugfactory.exportattribute
//
// 导出所有项目的药物属性和药品信息
type AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) Reset() {
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.exportattribute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest()
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryExportattributeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest
func GetAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest() *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest {
	return poolAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest.Get().(*AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest 将 AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest(v *AlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryExportattributeAPIRequest.Put(v)
}
