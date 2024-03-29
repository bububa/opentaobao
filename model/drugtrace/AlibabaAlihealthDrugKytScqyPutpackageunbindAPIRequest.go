package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest 码拼箱解除父子关系接口 API请求
// alibaba.alihealth.drug.kyt.scqy.putpackageunbind
//
// 码拼箱解除父子关系接口
type AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 解绑的父码
	_parentCode string
	// 解绑的子码
	_childCodes string
}

// NewAlibabaAlihealthDrugKytScqyPutpackageunbindRequest 初始化AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest对象
func NewAlibabaAlihealthDrugKytScqyPutpackageunbindRequest() *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest {
	return &AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) Reset() {
	r._refEntId = ""
	r._parentCode = ""
	r._childCodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.putpackageunbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetParentCode is ParentCode Setter
// 解绑的父码
func (r *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) SetParentCode(_parentCode string) error {
	r._parentCode = _parentCode
	r.Set("parent_code", _parentCode)
	return nil
}

// GetParentCode ParentCode Getter
func (r AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) GetParentCode() string {
	return r._parentCode
}

// SetChildCodes is ChildCodes Setter
// 解绑的子码
func (r *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) SetChildCodes(_childCodes string) error {
	r._childCodes = _childCodes
	r.Set("child_codes", _childCodes)
	return nil
}

// GetChildCodes ChildCodes Getter
func (r AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) GetChildCodes() string {
	return r._childCodes
}

var poolAlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytScqyPutpackageunbindRequest()
	},
}

// GetAlibabaAlihealthDrugKytScqyPutpackageunbindRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest
func GetAlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest() *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest {
	return poolAlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest.Get().(*AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest 将 AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest(v *AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest.Put(v)
}
