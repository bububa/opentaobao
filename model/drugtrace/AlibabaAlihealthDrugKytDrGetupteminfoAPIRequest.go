package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest 获取上游温度信息（疫苗） API请求
// alibaba.alihealth.drug.kyt.dr.getupteminfo
//
// 根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
type AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID
	_refEntId string
}

// NewAlibabaAlihealthDrugKytDrGetupteminfoRequest 初始化AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest对象
func NewAlibabaAlihealthDrugKytDrGetupteminfoRequest() *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest {
	return &AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) Reset() {
	r._code = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.getupteminfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytDrGetupteminfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrGetupteminfoRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrGetupteminfoRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest
func GetAlibabaAlihealthDrugKytDrGetupteminfoAPIRequest() *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest {
	return poolAlibabaAlihealthDrugKytDrGetupteminfoAPIRequest.Get().(*AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrGetupteminfoAPIRequest 将 AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrGetupteminfoAPIRequest(v *AlibabaAlihealthDrugKytDrGetupteminfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrGetupteminfoAPIRequest.Put(v)
}
