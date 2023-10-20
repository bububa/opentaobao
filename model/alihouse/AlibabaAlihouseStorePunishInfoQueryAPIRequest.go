package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousestorepunishinfoqueryAPIRequest 门店处罚信息查询 API请求
// alibaba.alihouse.store.punish.info.query
//
// 门店处罚信息查询
type AlibabaalihousestorepunishinfoqueryAPIRequest struct {
	model.Params
	// dto
	_queryStorePunishDto *QueryStorePunishDto
}

// NewAlibabaalihousestorepunishinfoqueryRequest 初始化AlibabaalihousestorepunishinfoqueryAPIRequest对象
func NewAlibabaalihousestorepunishinfoqueryRequest() *AlibabaalihousestorepunishinfoqueryAPIRequest {
	return &AlibabaalihousestorepunishinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousestorepunishinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.store.punish.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousestorepunishinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousestorepunishinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryStorePunishDto is QueryStorePunishDto Setter
// dto
func (r *AlibabaalihousestorepunishinfoqueryAPIRequest) SetQueryStorePunishDto(_queryStorePunishDto *QueryStorePunishDto) error {
	r._queryStorePunishDto = _queryStorePunishDto
	r.Set("query_store_punish_dto", _queryStorePunishDto)
	return nil
}

// GetQueryStorePunishDto QueryStorePunishDto Getter
func (r AlibabaalihousestorepunishinfoqueryAPIRequest) GetQueryStorePunishDto() *QueryStorePunishDto {
	return r._queryStorePunishDto
}
