package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSebpOrganizationGetinviteinfoAPIRequest 淘小铺机构上下级关系 API请求
// taobao.sebp.organization.getinviteinfo
//
// 机构人员获取机构上下级关系信息
type TaobaoSebpOrganizationGetinviteinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 第几页
	_pageNum int64
}

// NewTaobaoSebpOrganizationGetinviteinfoRequest 初始化TaobaoSebpOrganizationGetinviteinfoAPIRequest对象
func NewTaobaoSebpOrganizationGetinviteinfoRequest() *TaobaoSebpOrganizationGetinviteinfoAPIRequest {
	return &TaobaoSebpOrganizationGetinviteinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSebpOrganizationGetinviteinfoAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.organization.getinviteinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSebpOrganizationGetinviteinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSebpOrganizationGetinviteinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyDate is ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaoSebpOrganizationGetinviteinfoAPIRequest) SetModifyDate(_modifyDate string) error {
	r._modifyDate = _modifyDate
	r.Set("modify_date", _modifyDate)
	return nil
}

// GetModifyDate ModifyDate Getter
func (r TaobaoSebpOrganizationGetinviteinfoAPIRequest) GetModifyDate() string {
	return r._modifyDate
}

// SetPageNum is PageNum Setter
// 第几页
func (r *TaobaoSebpOrganizationGetinviteinfoAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoSebpOrganizationGetinviteinfoAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
