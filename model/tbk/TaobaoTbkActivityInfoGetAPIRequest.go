package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkactivityinfogetAPIRequest 淘宝客-推广者-官方活动转链 API请求
// taobao.tbk.activity.info.get
//
// 支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
type TaobaotbkactivityinfogetAPIRequest struct {
	model.Params
	// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
	_activityMaterialId string
	// mm_xxx_xxx_xxx 仅三方分成场景使用
	_subPid string
	// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_unionId string
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
	// 渠道关系id
	_relationId int64
}

// NewTaobaotbkactivityinfogetRequest 初始化TaobaotbkactivityinfogetAPIRequest对象
func NewTaobaotbkactivityinfogetRequest() *TaobaotbkactivityinfogetAPIRequest {
	return &TaobaotbkactivityinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkactivityinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.activity.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkactivityinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkactivityinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityMaterialId is ActivityMaterialId Setter
// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
func (r *TaobaotbkactivityinfogetAPIRequest) SetActivityMaterialId(_activityMaterialId string) error {
	r._activityMaterialId = _activityMaterialId
	r.Set("activity_material_id", _activityMaterialId)
	return nil
}

// GetActivityMaterialId ActivityMaterialId Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetActivityMaterialId() string {
	return r._activityMaterialId
}

// SetSubPid is SubPid Setter
// mm_xxx_xxx_xxx 仅三方分成场景使用
func (r *TaobaotbkactivityinfogetAPIRequest) SetSubPid(_subPid string) error {
	r._subPid = _subPid
	r.Set("sub_pid", _subPid)
	return nil
}

// GetSubPid SubPid Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetSubPid() string {
	return r._subPid
}

// SetUnionId is UnionId Setter
// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (r *TaobaotbkactivityinfogetAPIRequest) SetUnionId(_unionId string) error {
	r._unionId = _unionId
	r.Set("union_id", _unionId)
	return nil
}

// GetUnionId UnionId Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetUnionId() string {
	return r._unionId
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaotbkactivityinfogetAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetRelationId is RelationId Setter
// 渠道关系id
func (r *TaobaotbkactivityinfogetAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetRelationId() int64 {
	return r._relationId
}
