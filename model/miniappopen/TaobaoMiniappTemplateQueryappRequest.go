package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询实例化应用版本 APIRequest
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

func NewTaobaoMiniappTemplateQueryappRequest() *TaobaoMiniappTemplateQueryappRequest{
    return &TaobaoMiniappTemplateQueryappRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappTemplateQueryappRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.queryapp"
}

func (r TaobaoMiniappTemplateQueryappRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappTemplateQueryappRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoMiniappTemplateQueryappRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoMiniappTemplateQueryappRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoMiniappTemplateQueryappRequest) GetTemplateId() string {
    return r.templateId
}

func (r *TaobaoMiniappTemplateQueryappRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoMiniappTemplateQueryappRequest) GetPageNum() int64 {
    return r.pageNum
}

