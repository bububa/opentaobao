package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest 医院药品子码申请接口 API请求
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
type AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest struct {
	model.Params
	// 企业ID（ref_ent_id)
	_refEntId string
	// 父码
	_code string
	// 申请数量
	_amount int64
}

// NewAlibabaAlihealthDrugCodeKytYyApplycodeRequest 初始化AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytYyApplycodeRequest() *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r._amount = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.yy.applycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID（ref_ent_id)
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 父码
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) GetCode() string {
	return r._code
}

// SetAmount is Amount Setter
// 申请数量
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) GetAmount() int64 {
	return r._amount
}

var poolAlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytYyApplycodeRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytYyApplycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest
func GetAlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest() *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest 将 AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest(v *AlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytYyApplycodeAPIRequest.Put(v)
}
