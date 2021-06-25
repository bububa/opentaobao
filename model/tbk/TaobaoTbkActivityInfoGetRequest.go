package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-官方活动转链 APIRequest
taobao.tbk.activity.info.get

支持入参推广位和官方活动会场ID，获取活动信息和推广链接，包含推广长链接、短链接、淘口令、微信推广二维码地址等。改该API支持二方、三方类型的转链。官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取。
*/
type TaobaoTbkActivityInfoGetRequest struct {
    model.Params

    // mm_xxx_xxx_xxx的第三位
    adzoneId   int64 

    // mm_xxx_xxx_xxx 仅三方分成场景使用
    subPid   string 

    // 渠道关系id
    relationId   int64 

    // 官方活动会场ID，从淘宝客后台“我要推广-活动推广”中获取
    activityMaterialId   string 

    // 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
    unionId   string 

}

func NewTaobaoTbkActivityInfoGetRequest() *TaobaoTbkActivityInfoGetRequest{
    return &TaobaoTbkActivityInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkActivityInfoGetRequest) GetApiMethodName() string {
    return "taobao.tbk.activity.info.get"
}

func (r TaobaoTbkActivityInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkActivityInfoGetRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

func (r TaobaoTbkActivityInfoGetRequest) GetAdzoneId() int64 {
    return r.adzoneId
}

func (r *TaobaoTbkActivityInfoGetRequest) SetSubPid(subPid string) error {
    r.subPid = subPid
    r.Set("sub_pid", subPid)
    return nil
}

func (r TaobaoTbkActivityInfoGetRequest) GetSubPid() string {
    return r.subPid
}

func (r *TaobaoTbkActivityInfoGetRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

func (r TaobaoTbkActivityInfoGetRequest) GetRelationId() int64 {
    return r.relationId
}

func (r *TaobaoTbkActivityInfoGetRequest) SetActivityMaterialId(activityMaterialId string) error {
    r.activityMaterialId = activityMaterialId
    r.Set("activity_material_id", activityMaterialId)
    return nil
}

func (r TaobaoTbkActivityInfoGetRequest) GetActivityMaterialId() string {
    return r.activityMaterialId
}

func (r *TaobaoTbkActivityInfoGetRequest) SetUnionId(unionId string) error {
    r.unionId = unionId
    r.Set("union_id", unionId)
    return nil
}

func (r TaobaoTbkActivityInfoGetRequest) GetUnionId() string {
    return r.unionId
}

