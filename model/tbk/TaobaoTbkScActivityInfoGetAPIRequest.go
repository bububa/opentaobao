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
	_activitymaterialid string
	// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_unionid string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneid int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteid int64
	// 渠道关系id
	_relationid int64
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

// SetActivitymaterialid is Activitymaterialid Setter
// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
func (r *TaobaotbkscactivityinfogetAPIRequest) SetActivitymaterialid(_activitymaterialid string) error {
	r._activitymaterialid = _activitymaterialid
	r.Set("activity_material_id", _activitymaterialid)
	return nil
}

// GetActivitymaterialid Activitymaterialid Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetActivitymaterialid() string {
	return r._activitymaterialid
}

// SetUnionid is Unionid Setter
// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (r *TaobaotbkscactivityinfogetAPIRequest) SetUnionid(_unionid string) error {
	r._unionid = _unionid
	r.Set("union_id", _unionid)
	return nil
}

// GetUnionid Unionid Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetUnionid() string {
	return r._unionid
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkscactivityinfogetAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetSiteid is Siteid Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaotbkscactivityinfogetAPIRequest) SetSiteid(_siteid int64) error {
	r._siteid = _siteid
	r.Set("site_id", _siteid)
	return nil
}

// GetSiteid Siteid Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetSiteid() int64 {
	return r._siteid
}

// SetRelationid is Relationid Setter
// 渠道关系id
func (r *TaobaotbkscactivityinfogetAPIRequest) SetRelationid(_relationid int64) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkscactivityinfogetAPIRequest) GetRelationid() int64 {
	return r._relationid
}
