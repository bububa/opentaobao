package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosebporganizationgetinviteinfoAPIRequest 淘小铺机构上下级关系 API请求
// taobao.sebp.organization.getinviteinfo
//
// 机构人员获取机构上下级关系信息
type TaobaosebporganizationgetinviteinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 第几页
	_pageNum int64
}

// NewTaobaosebporganizationgetinviteinfoRequest 初始化TaobaosebporganizationgetinviteinfoAPIRequest对象
func NewTaobaosebporganizationgetinviteinfoRequest() *TaobaosebporganizationgetinviteinfoAPIRequest {
	return &TaobaosebporganizationgetinviteinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosebporganizationgetinviteinfoAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.organization.getinviteinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosebporganizationgetinviteinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosebporganizationgetinviteinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyDate is ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaosebporganizationgetinviteinfoAPIRequest) SetModifyDate(_modifyDate string) error {
	r._modifyDate = _modifyDate
	r.Set("modify_date", _modifyDate)
	return nil
}

// GetModifyDate ModifyDate Getter
func (r TaobaosebporganizationgetinviteinfoAPIRequest) GetModifyDate() string {
	return r._modifyDate
}

// SetPageNum is PageNum Setter
// 第几页
func (r *TaobaosebporganizationgetinviteinfoAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaosebporganizationgetinviteinfoAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
