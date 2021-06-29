package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户备案信息查询 API请求
taobao.tbk.sc.publisher.info.get

查询已生成的渠道id或会员运营id的相关信息。
*/
type TaobaoTbkScPublisherInfoGetRequest struct {
    model.Params
    // 类型，必选 1:渠道信息；2:会员信息
    _infoType   int64
    // 渠道独占 - 渠道关系ID
    _relationId   int64
    // 第几页
    _pageNo   int64
    // 每页大小
    _pageSize   int64
    // 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
    _relationApp   string
    // 会员独占 - 会员运营ID
    _specialId   string
    // 淘宝客外部用户标记，如自身系统账户ID；微信ID等
    _externalId   string
    // 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
    _externalType   int64
}

// 初始化TaobaoTbkScPublisherInfoGetRequest对象
func NewTaobaoTbkScPublisherInfoGetRequest() *TaobaoTbkScPublisherInfoGetRequest{
    return &TaobaoTbkScPublisherInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkScPublisherInfoGetRequest) GetApiMethodName() string {
    return "taobao.tbk.sc.publisher.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkScPublisherInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InfoType Setter
// 类型，必选 1:渠道信息；2:会员信息
func (r *TaobaoTbkScPublisherInfoGetRequest) SetInfoType(_infoType int64) error {
    r._infoType = _infoType
    r.Set("info_type", _infoType)
    return nil
}

// InfoType Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetInfoType() int64 {
    return r._infoType
}
// RelationId Setter
// 渠道独占 - 渠道关系ID
func (r *TaobaoTbkScPublisherInfoGetRequest) SetRelationId(_relationId int64) error {
    r._relationId = _relationId
    r.Set("relation_id", _relationId)
    return nil
}

// RelationId Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetRelationId() int64 {
    return r._relationId
}
// PageNo Setter
// 第几页
func (r *TaobaoTbkScPublisherInfoGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页大小
func (r *TaobaoTbkScPublisherInfoGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// RelationApp Setter
// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
func (r *TaobaoTbkScPublisherInfoGetRequest) SetRelationApp(_relationApp string) error {
    r._relationApp = _relationApp
    r.Set("relation_app", _relationApp)
    return nil
}

// RelationApp Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetRelationApp() string {
    return r._relationApp
}
// SpecialId Setter
// 会员独占 - 会员运营ID
func (r *TaobaoTbkScPublisherInfoGetRequest) SetSpecialId(_specialId string) error {
    r._specialId = _specialId
    r.Set("special_id", _specialId)
    return nil
}

// SpecialId Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetSpecialId() string {
    return r._specialId
}
// ExternalId Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaoTbkScPublisherInfoGetRequest) SetExternalId(_externalId string) error {
    r._externalId = _externalId
    r.Set("external_id", _externalId)
    return nil
}

// ExternalId Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetExternalId() string {
    return r._externalId
}
// ExternalType Setter
// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
func (r *TaobaoTbkScPublisherInfoGetRequest) SetExternalType(_externalType int64) error {
    r._externalType = _externalType
    r.Set("external_type", _externalType)
    return nil
}

// ExternalType Getter
func (r TaobaoTbkScPublisherInfoGetRequest) GetExternalType() int64 {
    return r._externalType
}
