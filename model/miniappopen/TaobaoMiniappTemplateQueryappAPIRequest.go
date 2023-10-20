package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiapptemplatequeryappAPIRequest 查询实例化应用版本 API请求
// taobao.miniapp.template.queryapp
//
// 根据模板id和商家信息，查询实例化小程序版本查询
type TaobaominiapptemplatequeryappAPIRequest struct {
	model.Params
	// 模板id
	_templateId string
	// 分页大小，最大50，按照小程序Id倒序
	_pageSize int64
	// 分页号,>=1
	_pageNum int64
}

// NewTaobaominiapptemplatequeryappRequest 初始化TaobaominiapptemplatequeryappAPIRequest对象
func NewTaobaominiapptemplatequeryappRequest() *TaobaominiapptemplatequeryappAPIRequest {
	return &TaobaominiapptemplatequeryappAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiapptemplatequeryappAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.queryapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiapptemplatequeryappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiapptemplatequeryappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaominiapptemplatequeryappAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaominiapptemplatequeryappAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetPageSize is PageSize Setter
// 分页大小，最大50，按照小程序Id倒序
func (r *TaobaominiapptemplatequeryappAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaominiapptemplatequeryappAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 分页号,&gt;=1
func (r *TaobaominiapptemplatequeryappAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaominiapptemplatequeryappAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
