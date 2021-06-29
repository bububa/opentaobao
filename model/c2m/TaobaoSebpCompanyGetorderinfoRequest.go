package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺公司订单信息 API请求
taobao.sebp.company.getorderinfo

淘小铺合作公司获取公司订单信息，用于公司结算使用
*/
type TaobaoSebpCompanyGetorderinfoRequest struct {
    model.Params
    // null-请求所有，20200616-请求2020年6月16号的变更信息
    modifyDate   string
    // 第几页
    pageNum   int64
}

// 初始化TaobaoSebpCompanyGetorderinfoRequest对象
func NewTaobaoSebpCompanyGetorderinfoRequest() *TaobaoSebpCompanyGetorderinfoRequest{
    return &TaobaoSebpCompanyGetorderinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpCompanyGetorderinfoRequest) GetApiMethodName() string {
    return "taobao.sebp.company.getorderinfo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpCompanyGetorderinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModifyDate Setter
// null-请求所有，20200616-请求2020年6月16号的变更信息
func (r *TaobaoSebpCompanyGetorderinfoRequest) SetModifyDate(modifyDate string) error {
    r.modifyDate = modifyDate
    r.Set("modify_date", modifyDate)
    return nil
}

// ModifyDate Getter
func (r TaobaoSebpCompanyGetorderinfoRequest) GetModifyDate() string {
    return r.modifyDate
}
// PageNum Setter
// 第几页
func (r *TaobaoSebpCompanyGetorderinfoRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoSebpCompanyGetorderinfoRequest) GetPageNum() int64 {
    return r.pageNum
}
