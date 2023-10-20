package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscactivityinfogetAPIRequest 淘宝客-服务商-官方活动转链 API请求
// taobao.tbk.sc.activity.info.get
//
// 服务商使用。支持入参推广者对应的推广位和官方活动会场ID，获取活动信息和推广者的推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
type TaobaotbkscactivityinfogetAPIRequest struct {
	model.Params
	// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
	_activityMaterialId string
	// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_unionId string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteId int64
	// 渠道关系id
	_relationId int64
}

// NewTaobaotbkscactivityinfogetRequest 初始化TaobaotbkscactivityinfogetAPIRequest对象
func NewTaobaotbkscactivityinfogetRequest() *TaobaotbkscactivityinfogetAPIRequest {
	return &TaobaotbkscactivityinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscactivityinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.activity.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscactivityinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscactivityinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityMaterialId is ActivityMaterialId Setter
// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
func (r *TaobaotbkscactivityinfogetAPIRequest) SetActivityMaterialId(_activityMaterialId string) error {
	r._activityMaterialId = _activityMaterialId
	r.Set("activity_material_id", _activityMaterialId)
	return nil
}

// GetActivityMaterialId ActivityMaterialId Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetActivityMaterialId() string {
	return r._activityMaterialId
}

// SetUnionId is UnionId Setter
// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (r *TaobaotbkscactivityinfogetAPIRequest) SetUnionId(_unionId string) error {
	r._unionId = _unionId
	r.Set("union_id", _unionId)
	return nil
}

// GetUnionId UnionId Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetUnionId() string {
	return r._unionId
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkscactivityinfogetAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetSiteId is SiteId Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaotbkscactivityinfogetAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetSiteId() int64 {
	return r._siteId
}

// SetRelationId is RelationId Setter
// 渠道关系id
func (r *TaobaotbkscactivityinfogetAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetRelationId() int64 {
	return r._relationId
}
