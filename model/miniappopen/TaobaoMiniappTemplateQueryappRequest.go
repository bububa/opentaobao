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
type TaobaoMiniappTemplateQueryappAPIRequest struct {
    model.Params
    // 分页大小，最大50，按照小程序Id倒序
    _pageSize   int64
    // 模板id
    _templateId   string
    // 分页号,>=1
    _pageNum   int64
}

// 初始化TaobaoMiniappTemplateQueryappAPIRequest对象
func NewTaobaoMiniappTemplateQueryappRequest() *TaobaoMiniappTemplateQueryappAPIRequest{
    return &TaobaoMiniappTemplateQueryappAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.template.queryapp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 分页大小，最大50，按照小程序Id倒序
func (r *TaobaoMiniappTemplateQueryappAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateQueryappAPIRequest) SetTemplateId(_templateId string) error {
    r._templateId = _templateId
    r.Set("template_id", _templateId)
    return nil
}

// TemplateId Getter
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetTemplateId() string {
    return r._templateId
}
// PageNum Setter
// 分页号,>=1
func (r *TaobaoMiniappTemplateQueryappAPIRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetPageNum() int64 {
    return r._pageNum
}
