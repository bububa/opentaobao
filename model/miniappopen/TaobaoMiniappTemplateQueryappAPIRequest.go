package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateQueryappAPIRequest 查询实例化应用版本 API请求
// taobao.miniapp.template.queryapp
//
// 根据模板id和商家信息，查询实例化小程序版本查询
type TaobaoMiniappTemplateQueryappAPIRequest struct {
	model.Params
	// 模板id
	_templateId string
	// 分页大小，最大50，按照小程序Id倒序
	_pageSize int64
	// 分页号,>=1
	_pageNum int64
}

// NewTaobaoMiniappTemplateQueryappRequest 初始化TaobaoMiniappTemplateQueryappAPIRequest对象
func NewTaobaoMiniappTemplateQueryappRequest() *TaobaoMiniappTemplateQueryappAPIRequest {
	return &TaobaoMiniappTemplateQueryappAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappTemplateQueryappAPIRequest) Reset() {
	r._templateId = ""
	r._pageSize = 0
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.template.queryapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 模板id
func (r *TaobaoMiniappTemplateQueryappAPIRequest) SetTemplateId(_templateId string) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetTemplateId() string {
	return r._templateId
}

// SetPageSize is PageSize Setter
// 分页大小，最大50，按照小程序Id倒序
func (r *TaobaoMiniappTemplateQueryappAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 分页号,&gt;=1
func (r *TaobaoMiniappTemplateQueryappAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoMiniappTemplateQueryappAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolTaobaoMiniappTemplateQueryappAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappTemplateQueryappRequest()
	},
}

// GetTaobaoMiniappTemplateQueryappRequest 从 sync.Pool 获取 TaobaoMiniappTemplateQueryappAPIRequest
func GetTaobaoMiniappTemplateQueryappAPIRequest() *TaobaoMiniappTemplateQueryappAPIRequest {
	return poolTaobaoMiniappTemplateQueryappAPIRequest.Get().(*TaobaoMiniappTemplateQueryappAPIRequest)
}

// ReleaseTaobaoMiniappTemplateQueryappAPIRequest 将 TaobaoMiniappTemplateQueryappAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappTemplateQueryappAPIRequest(v *TaobaoMiniappTemplateQueryappAPIRequest) {
	v.Reset()
	poolTaobaoMiniappTemplateQueryappAPIRequest.Put(v)
}
