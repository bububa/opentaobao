package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugLsydSaveentAPIRequest 零售药店往来单位新增 API请求
// alibaba.alihealth.drug.lsyd.saveent
//
// 新增往来单位企业记录
type AlibabaAlihealthDrugLsydSaveentAPIRequest struct {
	model.Params
	// 添加企业唯一标识
	_refEntId string
	// 新增企业信息
	_addEntReq *AddEntReqDto
	// 图片数据流。图片大小务必控制在2M以内
	_licPictureByte *model.File
}

// NewAlibabaAlihealthDrugLsydSaveentRequest 初始化AlibabaAlihealthDrugLsydSaveentAPIRequest对象
func NewAlibabaAlihealthDrugLsydSaveentRequest() *AlibabaAlihealthDrugLsydSaveentAPIRequest {
	return &AlibabaAlihealthDrugLsydSaveentAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugLsydSaveentAPIRequest) Reset() {
	r._refEntId = ""
	r._addEntReq = nil
	r._licPictureByte = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugLsydSaveentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.lsyd.saveent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugLsydSaveentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugLsydSaveentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 添加企业唯一标识
func (r *AlibabaAlihealthDrugLsydSaveentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugLsydSaveentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetAddEntReq is AddEntReq Setter
// 新增企业信息
func (r *AlibabaAlihealthDrugLsydSaveentAPIRequest) SetAddEntReq(_addEntReq *AddEntReqDto) error {
	r._addEntReq = _addEntReq
	r.Set("add_ent_req", _addEntReq)
	return nil
}

// GetAddEntReq AddEntReq Getter
func (r AlibabaAlihealthDrugLsydSaveentAPIRequest) GetAddEntReq() *AddEntReqDto {
	return r._addEntReq
}

// SetLicPictureByte is LicPictureByte Setter
// 图片数据流。图片大小务必控制在2M以内
func (r *AlibabaAlihealthDrugLsydSaveentAPIRequest) SetLicPictureByte(_licPictureByte *model.File) error {
	r._licPictureByte = _licPictureByte
	r.Set("lic_picture_byte", _licPictureByte)
	return nil
}

// GetLicPictureByte LicPictureByte Getter
func (r AlibabaAlihealthDrugLsydSaveentAPIRequest) GetLicPictureByte() *model.File {
	return r._licPictureByte
}

var poolAlibabaAlihealthDrugLsydSaveentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugLsydSaveentRequest()
	},
}

// GetAlibabaAlihealthDrugLsydSaveentRequest 从 sync.Pool 获取 AlibabaAlihealthDrugLsydSaveentAPIRequest
func GetAlibabaAlihealthDrugLsydSaveentAPIRequest() *AlibabaAlihealthDrugLsydSaveentAPIRequest {
	return poolAlibabaAlihealthDrugLsydSaveentAPIRequest.Get().(*AlibabaAlihealthDrugLsydSaveentAPIRequest)
}

// ReleaseAlibabaAlihealthDrugLsydSaveentAPIRequest 将 AlibabaAlihealthDrugLsydSaveentAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugLsydSaveentAPIRequest(v *AlibabaAlihealthDrugLsydSaveentAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugLsydSaveentAPIRequest.Put(v)
}
