package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugUploadExtinfoAPIRequest 中药饮片及器械对接 API请求
// alibaba.alihealth.drug.upload.extinfo
//
// 中药饮片及器械对接
type AlibabaAlihealthDrugUploadExtinfoAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugId string
	// 批次
	_batchNo string
	// 扩展信息
	_extInfoDto *ExtInfoDto
}

// NewAlibabaAlihealthDrugUploadExtinfoRequest 初始化AlibabaAlihealthDrugUploadExtinfoAPIRequest对象
func NewAlibabaAlihealthDrugUploadExtinfoRequest() *AlibabaAlihealthDrugUploadExtinfoAPIRequest {
	return &AlibabaAlihealthDrugUploadExtinfoAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugUploadExtinfoAPIRequest) Reset() {
	r._refEntId = ""
	r._drugId = ""
	r._batchNo = ""
	r._extInfoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.upload.extinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugUploadExtinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDrugId is DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugUploadExtinfoAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// GetDrugId DrugId Getter
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetDrugId() string {
	return r._drugId
}

// SetBatchNo is BatchNo Setter
// 批次
func (r *AlibabaAlihealthDrugUploadExtinfoAPIRequest) SetBatchNo(_batchNo string) error {
	r._batchNo = _batchNo
	r.Set("batch_no", _batchNo)
	return nil
}

// GetBatchNo BatchNo Getter
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetBatchNo() string {
	return r._batchNo
}

// SetExtInfoDto is ExtInfoDto Setter
// 扩展信息
func (r *AlibabaAlihealthDrugUploadExtinfoAPIRequest) SetExtInfoDto(_extInfoDto *ExtInfoDto) error {
	r._extInfoDto = _extInfoDto
	r.Set("ext_info_dto", _extInfoDto)
	return nil
}

// GetExtInfoDto ExtInfoDto Getter
func (r AlibabaAlihealthDrugUploadExtinfoAPIRequest) GetExtInfoDto() *ExtInfoDto {
	return r._extInfoDto
}

var poolAlibabaAlihealthDrugUploadExtinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugUploadExtinfoRequest()
	},
}

// GetAlibabaAlihealthDrugUploadExtinfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugUploadExtinfoAPIRequest
func GetAlibabaAlihealthDrugUploadExtinfoAPIRequest() *AlibabaAlihealthDrugUploadExtinfoAPIRequest {
	return poolAlibabaAlihealthDrugUploadExtinfoAPIRequest.Get().(*AlibabaAlihealthDrugUploadExtinfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugUploadExtinfoAPIRequest 将 AlibabaAlihealthDrugUploadExtinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugUploadExtinfoAPIRequest(v *AlibabaAlihealthDrugUploadExtinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugUploadExtinfoAPIRequest.Put(v)
}
