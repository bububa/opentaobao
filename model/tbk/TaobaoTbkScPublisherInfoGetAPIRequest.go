package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScPublisherInfoGetAPIRequest 淘宝客-公用-私域用户备案信息查询 API请求
// taobao.tbk.sc.publisher.info.get
//
// 查询已生成的渠道id或会员运营id的相关信息。
type TaobaoTbkScPublisherInfoGetAPIRequest struct {
	model.Params
	// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
	_relationApp string
	// 会员独占 - 会员运营ID
	_specialId string
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	_externalId string
	// 类型，必选 1:渠道信息；2:会员信息
	_infoType int64
	// 渠道独占 - 渠道关系ID
	_relationId int64
	// 第几页
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
	_externalType int64
}

// NewTaobaoTbkScPublisherInfoGetRequest 初始化TaobaoTbkScPublisherInfoGetAPIRequest对象
func NewTaobaoTbkScPublisherInfoGetRequest() *TaobaoTbkScPublisherInfoGetAPIRequest {
	return &TaobaoTbkScPublisherInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.publisher.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRelationApp is RelationApp Setter
// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetRelationApp(_relationApp string) error {
	r._relationApp = _relationApp
	r.Set("relation_app", _relationApp)
	return nil
}

// GetRelationApp RelationApp Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetRelationApp() string {
	return r._relationApp
}

// SetSpecialId is SpecialId Setter
// 会员独占 - 会员运营ID
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetExternalId is ExternalId Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetExternalId(_externalId string) error {
	r._externalId = _externalId
	r.Set("external_id", _externalId)
	return nil
}

// GetExternalId ExternalId Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetExternalId() string {
	return r._externalId
}

// SetInfoType is InfoType Setter
// 类型，必选 1:渠道信息；2:会员信息
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetInfoType(_infoType int64) error {
	r._infoType = _infoType
	r.Set("info_type", _infoType)
	return nil
}

// GetInfoType InfoType Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetInfoType() int64 {
	return r._infoType
}

// SetRelationId is RelationId Setter
// 渠道独占 - 渠道关系ID
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetPageNo is PageNo Setter
// 第几页
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetExternalType is ExternalType Setter
// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
func (r *TaobaoTbkScPublisherInfoGetAPIRequest) SetExternalType(_externalType int64) error {
	r._externalType = _externalType
	r.Set("external_type", _externalType)
	return nil
}

// GetExternalType ExternalType Getter
func (r TaobaoTbkScPublisherInfoGetAPIRequest) GetExternalType() int64 {
	return r._externalType
}
