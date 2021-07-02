package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSebpCompanyGetorderinfoAPIRequest 淘小铺公司订单信息 API请求
// taobao.sebp.company.getorderinfo
//
// 淘小铺合作公司获取公司订单信息，用于公司结算使用
type TaobaoSebpCompanyGetorderinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 第几页
	_pageNum int64
}

// NewTaobaoSebpCompanyGetorderinfoRequest 初始化TaobaoSebpCompanyGetorderinfoAPIRequest对象
func NewTaobaoSebpCompanyGetorderinfoRequest() *TaobaoSebpCompanyGetorderinfoAPIRequest {
	return &TaobaoSebpCompanyGetorderinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSebpCompanyGetorderinfoAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.company.getorderinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSebpCompanyGetorderinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetModifyDate is ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaoSebpCompanyGetorderinfoAPIRequest) SetModifyDate(_modifyDate string) error {
	r._modifyDate = _modifyDate
	r.Set("modify_date", _modifyDate)
	return nil
}

// GetModifyDate ModifyDate Getter
func (r TaobaoSebpCompanyGetorderinfoAPIRequest) GetModifyDate() string {
	return r._modifyDate
}

// SetPageNum is PageNum Setter
// 第几页
func (r *TaobaoSebpCompanyGetorderinfoAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoSebpCompanyGetorderinfoAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
