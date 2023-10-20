package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdCreateAPIRequest 淘宝客-服务商-创建人群标签 API请求
// taobao.tbk.sc.ucrowd.create
//
// 服务商使用。可为淘宝客创建会员人群标签，获得人群id，人群id可用于物料评估推荐等场景。
type TaobaoTbkScUcrowdCreateAPIRequest struct {
	model.Params
	// 人群外部编码，只能包含：字母、数字、下划线（以字母、数字开头），最大128字符
	_externalCrowdCode string
	// 人群名称，最大32字符
	_ucrowdName string
	// 人群描述，最大100字符
	_desc string
	// 人群类型，1:社群
	_ucrowdType int64
}

// NewTaobaoTbkScUcrowdCreateRequest 初始化TaobaoTbkScUcrowdCreateAPIRequest对象
func NewTaobaoTbkScUcrowdCreateRequest() *TaobaoTbkScUcrowdCreateAPIRequest {
	return &TaobaoTbkScUcrowdCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExternalCrowdCode is ExternalCrowdCode Setter
// 人群外部编码，只能包含：字母、数字、下划线（以字母、数字开头），最大128字符
func (r *TaobaoTbkScUcrowdCreateAPIRequest) SetExternalCrowdCode(_externalCrowdCode string) error {
	r._externalCrowdCode = _externalCrowdCode
	r.Set("external_crowd_code", _externalCrowdCode)
	return nil
}

// GetExternalCrowdCode ExternalCrowdCode Getter
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetExternalCrowdCode() string {
	return r._externalCrowdCode
}

// SetUcrowdName is UcrowdName Setter
// 人群名称，最大32字符
func (r *TaobaoTbkScUcrowdCreateAPIRequest) SetUcrowdName(_ucrowdName string) error {
	r._ucrowdName = _ucrowdName
	r.Set("ucrowd_name", _ucrowdName)
	return nil
}

// GetUcrowdName UcrowdName Getter
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetUcrowdName() string {
	return r._ucrowdName
}

// SetDesc is Desc Setter
// 人群描述，最大100字符
func (r *TaobaoTbkScUcrowdCreateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetDesc() string {
	return r._desc
}

// SetUcrowdType is UcrowdType Setter
// 人群类型，1:社群
func (r *TaobaoTbkScUcrowdCreateAPIRequest) SetUcrowdType(_ucrowdType int64) error {
	r._ucrowdType = _ucrowdType
	r.Set("ucrowd_type", _ucrowdType)
	return nil
}

// GetUcrowdType UcrowdType Getter
func (r TaobaoTbkScUcrowdCreateAPIRequest) GetUcrowdType() int64 {
	return r._ucrowdType
}
