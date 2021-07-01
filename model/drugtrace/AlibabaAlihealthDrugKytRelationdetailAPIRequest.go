package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytRelationdetailAPIRequest
关联关系处理详情 API请求
alibaba.alihealth.drug.kyt.relationdetail

关联关系处理详情 */
type AlibabaAlihealthDrugKytRelationdetailAPIRequest struct {
	model.Params
	// 码激活文件上传信息标识
	_codeActiveInfoId string
	// 企业ID
	_refEntId string
	// 客户端ID【默认写2】
	_clientType string
}

// NewAlibabaAlihealthDrugKytRelationdetailRequest 初始化AlibabaAlihealthDrugKytRelationdetailAPIRequest对象
func NewAlibabaAlihealthDrugKytRelationdetailRequest() *AlibabaAlihealthDrugKytRelationdetailAPIRequest {
	return &AlibabaAlihealthDrugKytRelationdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytRelationdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.relationdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytRelationdetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CodeActiveInfoId Setter
// 码激活文件上传信息标识
func (r *AlibabaAlihealthDrugKytRelationdetailAPIRequest) SetCodeActiveInfoId(_codeActiveInfoId string) error {
	r._codeActiveInfoId = _codeActiveInfoId
	r.Set("code_active_info_id", _codeActiveInfoId)
	return nil
}

// Get CodeActiveInfoId Getter
func (r AlibabaAlihealthDrugKytRelationdetailAPIRequest) GetCodeActiveInfoId() string {
	return r._codeActiveInfoId
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytRelationdetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytRelationdetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is ClientType Setter
// 客户端ID【默认写2】
func (r *AlibabaAlihealthDrugKytRelationdetailAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// Get ClientType Getter
func (r AlibabaAlihealthDrugKytRelationdetailAPIRequest) GetClientType() string {
	return r._clientType
}
