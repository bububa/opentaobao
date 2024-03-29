package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSaveentAPIRequest 新增往来单位企业 API请求
// alibaba.alihealth.drug.kyt.saveent
//
// 新增往来单位企业记录
type AlibabaAlihealthDrugKytSaveentAPIRequest struct {
	model.Params
	// 添加企业唯一标识
	_refEntId string
	// 新增企业信息
	_addEntReq *AddEntReqDto
	// 图片数据流。图片大小务必控制在2M以内
	_licPictureByte *model.File
}

// NewAlibabaAlihealthDrugKytSaveentRequest 初始化AlibabaAlihealthDrugKytSaveentAPIRequest对象
func NewAlibabaAlihealthDrugKytSaveentRequest() *AlibabaAlihealthDrugKytSaveentAPIRequest {
	return &AlibabaAlihealthDrugKytSaveentAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) Reset() {
	r._refEntId = ""
	r._addEntReq = nil
	r._licPictureByte = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.saveent"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 添加企业唯一标识
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetAddEntReq is AddEntReq Setter
// 新增企业信息
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) SetAddEntReq(_addEntReq *AddEntReqDto) error {
	r._addEntReq = _addEntReq
	r.Set("add_ent_req", _addEntReq)
	return nil
}

// GetAddEntReq AddEntReq Getter
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetAddEntReq() *AddEntReqDto {
	return r._addEntReq
}

// SetLicPictureByte is LicPictureByte Setter
// 图片数据流。图片大小务必控制在2M以内
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) SetLicPictureByte(_licPictureByte *model.File) error {
	r._licPictureByte = _licPictureByte
	r.Set("lic_picture_byte", _licPictureByte)
	return nil
}

// GetLicPictureByte LicPictureByte Getter
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetLicPictureByte() *model.File {
	return r._licPictureByte
}

var poolAlibabaAlihealthDrugKytSaveentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytSaveentRequest()
	},
}

// GetAlibabaAlihealthDrugKytSaveentRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytSaveentAPIRequest
func GetAlibabaAlihealthDrugKytSaveentAPIRequest() *AlibabaAlihealthDrugKytSaveentAPIRequest {
	return poolAlibabaAlihealthDrugKytSaveentAPIRequest.Get().(*AlibabaAlihealthDrugKytSaveentAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytSaveentAPIRequest 将 AlibabaAlihealthDrugKytSaveentAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSaveentAPIRequest(v *AlibabaAlihealthDrugKytSaveentAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSaveentAPIRequest.Put(v)
}
