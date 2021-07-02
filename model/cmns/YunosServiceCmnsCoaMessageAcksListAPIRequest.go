package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageAcksListAPIRequest 消息ack记录查询 API请求
// yunos.service.cmns.coa.message.acks.list
//
// 第三方应用开发者调用此接口查询消息ack记录
type YunosServiceCmnsCoaMessageAcksListAPIRequest struct {
	model.Params
	// 消息id
	_mid int64
	// 设备id
	_did int64
	// 分页查询页码
	_pageIndex int64
	// 分页每页数据集数
	_pageSize int64
}

// NewYunosServiceCmnsCoaMessageAcksListRequest 初始化YunosServiceCmnsCoaMessageAcksListAPIRequest对象
func NewYunosServiceCmnsCoaMessageAcksListRequest() *YunosServiceCmnsCoaMessageAcksListAPIRequest {
	return &YunosServiceCmnsCoaMessageAcksListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageAcksListAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.message.acks.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageAcksListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMid is Mid Setter
// 消息id
func (r *YunosServiceCmnsCoaMessageAcksListAPIRequest) SetMid(_mid int64) error {
	r._mid = _mid
	r.Set("mid", _mid)
	return nil
}

// GetMid Mid Getter
func (r YunosServiceCmnsCoaMessageAcksListAPIRequest) GetMid() int64 {
	return r._mid
}

// SetDid is Did Setter
// 设备id
func (r *YunosServiceCmnsCoaMessageAcksListAPIRequest) SetDid(_did int64) error {
	r._did = _did
	r.Set("did", _did)
	return nil
}

// GetDid Did Getter
func (r YunosServiceCmnsCoaMessageAcksListAPIRequest) GetDid() int64 {
	return r._did
}

// SetPageIndex is PageIndex Setter
// 分页查询页码
func (r *YunosServiceCmnsCoaMessageAcksListAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r YunosServiceCmnsCoaMessageAcksListAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页每页数据集数
func (r *YunosServiceCmnsCoaMessageAcksListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosServiceCmnsCoaMessageAcksListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
