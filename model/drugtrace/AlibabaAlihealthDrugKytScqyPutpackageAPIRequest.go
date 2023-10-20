package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyPutpackageAPIRequest 码拼箱 API请求
// alibaba.alihealth.drug.kyt.scqy.putpackage
//
// 码拼箱接口
type AlibabaAlihealthDrugKytScqyPutpackageAPIRequest struct {
	model.Params
	// 企业id
	_refEntid string
	// 二级码
	_secondaryCode string
	// 一级码
	_primaryCodes string
}

// NewAlibabaAlihealthDrugKytScqyPutpackageRequest 初始化AlibabaAlihealthDrugKytScqyPutpackageAPIRequest对象
func NewAlibabaAlihealthDrugKytScqyPutpackageRequest() *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest {
	return &AlibabaAlihealthDrugKytScqyPutpackageAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) Reset() {
	r._refEntid = ""
	r._secondaryCode = ""
	r._primaryCodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.putpackage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntid is RefEntid Setter
// 企业id
func (r *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) SetRefEntid(_refEntid string) error {
	r._refEntid = _refEntid
	r.Set("ref_entid", _refEntid)
	return nil
}

// GetRefEntid RefEntid Getter
func (r AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) GetRefEntid() string {
	return r._refEntid
}

// SetSecondaryCode is SecondaryCode Setter
// 二级码
func (r *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) SetSecondaryCode(_secondaryCode string) error {
	r._secondaryCode = _secondaryCode
	r.Set("secondary_code", _secondaryCode)
	return nil
}

// GetSecondaryCode SecondaryCode Getter
func (r AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) GetSecondaryCode() string {
	return r._secondaryCode
}

// SetPrimaryCodes is PrimaryCodes Setter
// 一级码
func (r *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) SetPrimaryCodes(_primaryCodes string) error {
	r._primaryCodes = _primaryCodes
	r.Set("primary_codes", _primaryCodes)
	return nil
}

// GetPrimaryCodes PrimaryCodes Getter
func (r AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) GetPrimaryCodes() string {
	return r._primaryCodes
}

var poolAlibabaAlihealthDrugKytScqyPutpackageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytScqyPutpackageRequest()
	},
}

// GetAlibabaAlihealthDrugKytScqyPutpackageRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyPutpackageAPIRequest
func GetAlibabaAlihealthDrugKytScqyPutpackageAPIRequest() *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest {
	return poolAlibabaAlihealthDrugKytScqyPutpackageAPIRequest.Get().(*AlibabaAlihealthDrugKytScqyPutpackageAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytScqyPutpackageAPIRequest 将 AlibabaAlihealthDrugKytScqyPutpackageAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyPutpackageAPIRequest(v *AlibabaAlihealthDrugKytScqyPutpackageAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyPutpackageAPIRequest.Put(v)
}
