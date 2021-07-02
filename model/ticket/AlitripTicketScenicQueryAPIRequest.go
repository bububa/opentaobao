package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketScenicQueryAPIRequest 【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询） API请求
// alitrip.ticket.scenic.query
//
// 查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
type AlitripTicketScenicQueryAPIRequest struct {
	model.Params
	// 标准景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
	_aliScenicId int64
	// 当前分页。每页默认最多返回20条数据
	_currentPage int64
	// 商家景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
	_outScenicId string
}

// NewAlitripTicketScenicQueryRequest 初始化AlitripTicketScenicQueryAPIRequest对象
func NewAlitripTicketScenicQueryRequest() *AlitripTicketScenicQueryAPIRequest {
	return &AlitripTicketScenicQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketScenicQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.scenic.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketScenicQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AliScenicId Setter
// 标准景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
func (r *AlitripTicketScenicQueryAPIRequest) SetAliScenicId(_aliScenicId int64) error {
	r._aliScenicId = _aliScenicId
	r.Set("ali_scenic_id", _aliScenicId)
	return nil
}

// Get AliScenicId Getter
func (r AlitripTicketScenicQueryAPIRequest) GetAliScenicId() int64 {
	return r._aliScenicId
}

// Set is CurrentPage Setter
// 当前分页。每页默认最多返回20条数据
func (r *AlitripTicketScenicQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// Get CurrentPage Getter
func (r AlitripTicketScenicQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// Set is OutScenicId Setter
// 商家景点ID。ali_scenic_id，out_scenic_id二者至少需要填写一个
func (r *AlitripTicketScenicQueryAPIRequest) SetOutScenicId(_outScenicId string) error {
	r._outScenicId = _outScenicId
	r.Set("out_scenic_id", _outScenicId)
	return nil
}

// Get OutScenicId Getter
func (r AlitripTicketScenicQueryAPIRequest) GetOutScenicId() string {
	return r._outScenicId
}
