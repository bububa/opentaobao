package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询实例化应用版本 API请求
taobao.miniapp.template.queryapp

根据模板id和商家信息，查询实例化小程序版本查询
*/
type TaobaoMiniappTemplateQueryappRequest struct {
    model.Params
    // 分页大小，最大50，按照小程序Id倒序
    pageSize   int64
    // 模板id
    templateId   string
    // 分页号,>=1
    pageNum   int64
}

// 初始化TaobaoMiniappTemplateQueryappRequest对象
func NewTaobaoMiniappTemplateQueryappRequest() *TaobaoMiniappTemplateQueryappRequest{
    return &TaobaoMiniappTemplateQueryappRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateQueryappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.queryapp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateQueryappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 分页大小，最大50，按照小程序Id倒序
func (r *TaobaoMiniappTemplateQueryappRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoMiniappTemplateQueryappRequest) GetPageSize() int64 {
    return r.pageSize
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateQueryappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateQueryappRequest) GetTemplateId() string {
    return r.templateId
}
// PageNum Setter
// 分页号,>=1
func (r *TaobaoMiniappTemplateQueryappRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoMiniappTemplateQueryappRequest) GetPageNum() int64 {
    return r.pageNum
}
