package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest 医院根据码查询码信息 API请求
// alibaba.alihealth.drug.code.kyt.yy.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 企业唯一标识（或appkey）
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytYyQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytYyQuerycodeRequest() *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.yy.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytYyQuerycodeRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytYyQuerycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest
func GetAlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest() *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest 将 AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest(v *AlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytYyQuerycodeAPIRequest.Put(v)
}
