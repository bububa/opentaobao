package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowdcreateAPIRequest 淘宝客-服务商-创建人群标签 API请求
// taobao.tbk.sc.ucrowd.create
//
// 服务商使用。可为淘宝客创建会员人群标签，获得人群id，人群id可用于物料评估推荐等场景。
type TaobaotbkscucrowdcreateAPIRequest struct {
	model.Params
	// 人群外部编码，只能包含：字母、数字、下划线（以字母、数字开头），最大128字符
	_externalcrowdcode string
	// 人群名称，最大32字符
	_ucrowdname string
	// 人群描述，最大100字符
	_desc string
	// 人群类型，1:社群
	_ucrowdtype int64
}

// NewTaobaotbkscucrowdcreateRequest 初始化TaobaotbkscucrowdcreateAPIRequest对象
func NewTaobaotbkscucrowdcreateRequest() *TaobaotbkscucrowdcreateAPIRequest {
	return &TaobaotbkscucrowdcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscucrowdcreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscucrowdcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscucrowdcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExternalcrowdcode is Externalcrowdcode Setter
// 人群外部编码，只能包含：字母、数字、下划线（以字母、数字开头），最大128字符
func (r *TaobaotbkscucrowdcreateAPIRequest) SetExternalcrowdcode(_externalcrowdcode string) error {
	r._externalcrowdcode = _externalcrowdcode
	r.Set("external_crowd_code", _externalcrowdcode)
	return nil
}

// GetExternalcrowdcode Externalcrowdcode Getter
func (r TaobaotbkscucrowdcreateAPIRequest) GetExternalcrowdcode() string {
	return r._externalcrowdcode
}

// SetUcrowdname is Ucrowdname Setter
// 人群名称，最大32字符
func (r *TaobaotbkscucrowdcreateAPIRequest) SetUcrowdname(_ucrowdname string) error {
	r._ucrowdname = _ucrowdname
	r.Set("ucrowd_name", _ucrowdname)
	return nil
}

// GetUcrowdname Ucrowdname Getter
func (r TaobaotbkscucrowdcreateAPIRequest) GetUcrowdname() string {
	return r._ucrowdname
}

// SetDesc is Desc Setter
// 人群描述，最大100字符
func (r *TaobaotbkscucrowdcreateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaotbkscucrowdcreateAPIRequest) GetDesc() string {
	return r._desc
}

// SetUcrowdtype is Ucrowdtype Setter
// 人群类型，1:社群
func (r *TaobaotbkscucrowdcreateAPIRequest) SetUcrowdtype(_ucrowdtype int64) error {
	r._ucrowdtype = _ucrowdtype
	r.Set("ucrowd_type", _ucrowdtype)
	return nil
}

// GetUcrowdtype Ucrowdtype Getter
func (r TaobaotbkscucrowdcreateAPIRequest) GetUcrowdtype() int64 {
	return r._ucrowdtype
}
