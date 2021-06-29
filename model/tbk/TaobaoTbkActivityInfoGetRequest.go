package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-官方活动转链 API请求
taobao.tbk.activity.info.get

支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
*/
type TaobaoTbkActivityInfoGetRequest struct {
    model.Params
    // mm_xxx_xxx_xxx的第三位
    _adzoneId   int64
    // mm_xxx_xxx_xxx 仅三方分成场景使用
    _subPid   string
    // 渠道关系id
    _relationId   int64
    // 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
    _activityMaterialId   string
    // 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
    _unionId   string
}

// 初始化TaobaoTbkActivityInfoGetRequest对象
func NewTaobaoTbkActivityInfoGetRequest() *TaobaoTbkActivityInfoGetRequest{
    return &TaobaoTbkActivityInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkActivityInfoGetRequest) GetApiMethodName() string {
    return "taobao.tbk.activity.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkActivityInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaoTbkActivityInfoGetRequest) SetAdzoneId(_adzoneId int64) error {
    r._adzoneId = _adzoneId
    r.Set("adzone_id", _adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkActivityInfoGetRequest) GetAdzoneId() int64 {
    return r._adzoneId
}
// SubPid Setter
// mm_xxx_xxx_xxx 仅三方分成场景使用
func (r *TaobaoTbkActivityInfoGetRequest) SetSubPid(_subPid string) error {
    r._subPid = _subPid
    r.Set("sub_pid", _subPid)
    return nil
}

// SubPid Getter
func (r TaobaoTbkActivityInfoGetRequest) GetSubPid() string {
    return r._subPid
}
// RelationId Setter
// 渠道关系id
func (r *TaobaoTbkActivityInfoGetRequest) SetRelationId(_relationId int64) error {
    r._relationId = _relationId
    r.Set("relation_id", _relationId)
    return nil
}

// RelationId Getter
func (r TaobaoTbkActivityInfoGetRequest) GetRelationId() int64 {
    return r._relationId
}
// ActivityMaterialId Setter
// 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
func (r *TaobaoTbkActivityInfoGetRequest) SetActivityMaterialId(_activityMaterialId string) error {
    r._activityMaterialId = _activityMaterialId
    r.Set("activity_material_id", _activityMaterialId)
    return nil
}

// ActivityMaterialId Getter
func (r TaobaoTbkActivityInfoGetRequest) GetActivityMaterialId() string {
    return r._activityMaterialId
}
// UnionId Setter
// 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (r *TaobaoTbkActivityInfoGetRequest) SetUnionId(_unionId string) error {
    r._unionId = _unionId
    r.Set("union_id", _unionId)
    return nil
}

// UnionId Getter
func (r TaobaoTbkActivityInfoGetRequest) GetUnionId() string {
    return r._unionId
}
