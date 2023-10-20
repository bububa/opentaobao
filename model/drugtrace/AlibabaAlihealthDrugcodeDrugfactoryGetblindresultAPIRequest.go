package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest 获取盲底文件处理结果 API请求
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 盲底文件名称
	_blindFileName string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) Reset() {
	r._refEntId = ""
	r._blindFileName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.getblindresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBlindFileName is BlindFileName Setter
// 盲底文件名称
func (r *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) SetBlindFileName(_blindFileName string) error {
	r._blindFileName = _blindFileName
	r.Set("blind_file_name", _blindFileName)
	return nil
}

// GetBlindFileName BlindFileName Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) GetBlindFileName() string {
	return r._blindFileName
}

var poolAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest()
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryGetblindresultRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest
func GetAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest() *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest {
	return poolAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest.Get().(*AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest 将 AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest(v *AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest.Put(v)
}
