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
	_activitymaterialid string
	// mm_xxx_xxx_xxx 仅三方分成场景使用
	_subpid string
	// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_unionid string
	// mm_xxx_xxx_xxx的第三位
	_adzoneid int64
	// 渠道关系id
	_relationid int64
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

// SetActivitymaterialid is Activitymaterialid Setter
// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
func (r *TaobaotbkactivityinfogetAPIRequest) SetActivitymaterialid(_activitymaterialid string) error {
	r._activitymaterialid = _activitymaterialid
	r.Set("activity_material_id", _activitymaterialid)
	return nil
}

// GetActivitymaterialid Activitymaterialid Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetActivitymaterialid() string {
	return r._activitymaterialid
}

// SetSubpid is Subpid Setter
// mm_xxx_xxx_xxx 仅三方分成场景使用
func (r *TaobaotbkactivityinfogetAPIRequest) SetSubpid(_subpid string) error {
	r._subpid = _subpid
	r.Set("sub_pid", _subpid)
	return nil
}

// GetSubpid Subpid Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetSubpid() string {
	return r._subpid
}

// SetUnionid is Unionid Setter
// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (r *TaobaotbkactivityinfogetAPIRequest) SetUnionid(_unionid string) error {
	r._unionid = _unionid
	r.Set("union_id", _unionid)
	return nil
}

// GetUnionid Unionid Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetUnionid() string {
	return r._unionid
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaotbkactivityinfogetAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}

// SetRelationid is Relationid Setter
// 渠道关系id
func (r *TaobaotbkactivityinfogetAPIRequest) SetRelationid(_relationid int64) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkactivityinfogetAPIRequest) GetRelationid() int64 {
	return r._relationid
}
