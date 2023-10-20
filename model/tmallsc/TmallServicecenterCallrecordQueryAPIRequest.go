package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterCallrecordQueryAPIRequest 天猫服务平台服务商查询通话记录接口 API请求
// tmall.servicecenter.callrecord.query
//
// 天猫服务平台服务商查询通话记录接口
type TmallServicecenterCallrecordQueryAPIRequest struct {
	model.Params
	// 类型
	_bizIdentifyType string
	// 工单id
	_bizIdentify string
	// 当前页数
	_pageNum int64
	// 每页查询数量
	_pageSize int64
}

// NewTmallServicecenterCallrecordQueryRequest 初始化TmallServicecenterCallrecordQueryAPIRequest对象
func NewTmallServicecenterCallrecordQueryRequest() *TmallServicecenterCallrecordQueryAPIRequest {
	return &TmallServicecenterCallrecordQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterCallrecordQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.callrecord.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterCallrecordQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterCallrecordQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizIdentifyType is BizIdentifyType Setter
// 类型
func (r *TmallServicecenterCallrecordQueryAPIRequest) SetBizIdentifyType(_bizIdentifyType string) error {
	r._bizIdentifyType = _bizIdentifyType
	r.Set("biz_identify_type", _bizIdentifyType)
	return nil
}

// GetBizIdentifyType BizIdentifyType Getter
func (r TmallServicecenterCallrecordQueryAPIRequest) GetBizIdentifyType() string {
	return r._bizIdentifyType
}

// SetBizIdentify is BizIdentify Setter
// 工单id
func (r *TmallServicecenterCallrecordQueryAPIRequest) SetBizIdentify(_bizIdentify string) error {
	r._bizIdentify = _bizIdentify
	r.Set("biz_identify", _bizIdentify)
	return nil
}

// GetBizIdentify BizIdentify Getter
func (r TmallServicecenterCallrecordQueryAPIRequest) GetBizIdentify() string {
	return r._bizIdentify
}

// SetPageNum is PageNum Setter
// 当前页数
func (r *TmallServicecenterCallrecordQueryAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TmallServicecenterCallrecordQueryAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 每页查询数量
func (r *TmallServicecenterCallrecordQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallServicecenterCallrecordQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
