package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScMembergroupOptionalAPIRequest 工具服务商member组查询、新增接口 API请求
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
type TaobaoTbkScMembergroupOptionalAPIRequest struct {
	model.Params
	// 淘宝数字id
	_tbNumIds string
	// member组id
	_memberGroupId int64
}

// NewTaobaoTbkScMembergroupOptionalRequest 初始化TaobaoTbkScMembergroupOptionalAPIRequest对象
func NewTaobaoTbkScMembergroupOptionalRequest() *TaobaoTbkScMembergroupOptionalAPIRequest {
	return &TaobaoTbkScMembergroupOptionalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScMembergroupOptionalAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.membergroup.optional"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScMembergroupOptionalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScMembergroupOptionalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbNumIds is TbNumIds Setter
// 淘宝数字id
func (r *TaobaoTbkScMembergroupOptionalAPIRequest) SetTbNumIds(_tbNumIds string) error {
	r._tbNumIds = _tbNumIds
	r.Set("tb_num_ids", _tbNumIds)
	return nil
}

// GetTbNumIds TbNumIds Getter
func (r TaobaoTbkScMembergroupOptionalAPIRequest) GetTbNumIds() string {
	return r._tbNumIds
}

// SetMemberGroupId is MemberGroupId Setter
// member组id
func (r *TaobaoTbkScMembergroupOptionalAPIRequest) SetMemberGroupId(_memberGroupId int64) error {
	r._memberGroupId = _memberGroupId
	r.Set("member_group_id", _memberGroupId)
	return nil
}

// GetMemberGroupId MemberGroupId Getter
func (r TaobaoTbkScMembergroupOptionalAPIRequest) GetMemberGroupId() int64 {
	return r._memberGroupId
}
