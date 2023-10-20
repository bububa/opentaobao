package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetbyentidAPIRequest 根据企业主键查看企业详细信息 API请求
// alibaba.alihealth.drug.kyt.getbyentid
//
// 根据企业主键查看企业详细信息
type AlibabaalihealthdrugkytgetbyentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业ID（返回该企业ID的详细信息）
	_entId string
}

// NewAlibabaalihealthdrugkytgetbyentidRequest 初始化AlibabaalihealthdrugkytgetbyentidAPIRequest对象
func NewAlibabaalihealthdrugkytgetbyentidRequest() *AlibabaalihealthdrugkytgetbyentidAPIRequest {
	return &AlibabaalihealthdrugkytgetbyentidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytgetbyentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getbyentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytgetbyentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytgetbyentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaalihealthdrugkytgetbyentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytgetbyentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntId is EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaalihealthdrugkytgetbyentidAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaalihealthdrugkytgetbyentidAPIRequest) GetEntId() string {
	return r._entId
}
