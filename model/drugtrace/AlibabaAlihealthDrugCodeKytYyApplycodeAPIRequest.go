package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytyyapplycodeAPIRequest 医院药品子码申请接口 API请求
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
type AlibabaalihealthdrugcodekytyyapplycodeAPIRequest struct {
	model.Params
	// 企业ID（ref_ent_id)
	_refEntId string
	// 父码
	_code string
	// 申请数量
	_amount int64
}

// NewAlibabaalihealthdrugcodekytyyapplycodeRequest 初始化AlibabaalihealthdrugcodekytyyapplycodeAPIRequest对象
func NewAlibabaalihealthdrugcodekytyyapplycodeRequest() *AlibabaalihealthdrugcodekytyyapplycodeAPIRequest {
	return &AlibabaalihealthdrugcodekytyyapplycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.yy.applycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID（ref_ent_id)
func (r *AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 父码
func (r *AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) GetCode() string {
	return r._code
}

// SetAmount is Amount Setter
// 申请数量
func (r *AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) SetAmount(_amount int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r AlibabaalihealthdrugcodekytyyapplycodeAPIRequest) GetAmount() int64 {
	return r._amount
}
