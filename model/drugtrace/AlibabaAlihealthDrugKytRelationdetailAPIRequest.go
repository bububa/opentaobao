package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytrelationdetailAPIRequest 关联关系处理详情 API请求
// alibaba.alihealth.drug.kyt.relationdetail
//
// 关联关系处理详情
type AlibabaalihealthdrugkytrelationdetailAPIRequest struct {
	model.Params
	// 码激活文件上传信息标识
	_codeActiveInfoId string
	// 企业ID
	_refEntId string
	// 客户端ID【默认写2】
	_clientType string
}

// NewAlibabaalihealthdrugkytrelationdetailRequest 初始化AlibabaalihealthdrugkytrelationdetailAPIRequest对象
func NewAlibabaalihealthdrugkytrelationdetailRequest() *AlibabaalihealthdrugkytrelationdetailAPIRequest {
	return &AlibabaalihealthdrugkytrelationdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytrelationdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.relationdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytrelationdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytrelationdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodeActiveInfoId is CodeActiveInfoId Setter
// 码激活文件上传信息标识
func (r *AlibabaalihealthdrugkytrelationdetailAPIRequest) SetCodeActiveInfoId(_codeActiveInfoId string) error {
	r._codeActiveInfoId = _codeActiveInfoId
	r.Set("code_active_info_id", _codeActiveInfoId)
	return nil
}

// GetCodeActiveInfoId CodeActiveInfoId Getter
func (r AlibabaalihealthdrugkytrelationdetailAPIRequest) GetCodeActiveInfoId() string {
	return r._codeActiveInfoId
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytrelationdetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytrelationdetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetClientType is ClientType Setter
// 客户端ID【默认写2】
func (r *AlibabaalihealthdrugkytrelationdetailAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaalihealthdrugkytrelationdetailAPIRequest) GetClientType() string {
	return r._clientType
}
