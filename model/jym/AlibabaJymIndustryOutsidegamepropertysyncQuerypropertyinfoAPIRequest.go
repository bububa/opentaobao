package jym

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest 外部查询游戏属性库属性信息 API请求
// alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo
//
// 外部查询游戏属性库属性信息
type AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest struct {
	model.Params
	// 属性信息查询DTO
	_outSideQueryGamePropertyInfoRequestDto *OutSideQueryGamePropertyInfoRequestDto
}

// NewAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoRequest 初始化AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest对象
func NewAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoRequest() *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest {
	return &AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) Reset() {
	r._outSideQueryGamePropertyInfoRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSideQueryGamePropertyInfoRequestDto is OutSideQueryGamePropertyInfoRequestDto Setter
// 属性信息查询DTO
func (r *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) SetOutSideQueryGamePropertyInfoRequestDto(_outSideQueryGamePropertyInfoRequestDto *OutSideQueryGamePropertyInfoRequestDto) error {
	r._outSideQueryGamePropertyInfoRequestDto = _outSideQueryGamePropertyInfoRequestDto
	r.Set("out_side_query_game_property_info_request_dto", _outSideQueryGamePropertyInfoRequestDto)
	return nil
}

// GetOutSideQueryGamePropertyInfoRequestDto OutSideQueryGamePropertyInfoRequestDto Getter
func (r AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) GetOutSideQueryGamePropertyInfoRequestDto() *OutSideQueryGamePropertyInfoRequestDto {
	return r._outSideQueryGamePropertyInfoRequestDto
}

var poolAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoRequest()
	},
}

// GetAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoRequest 从 sync.Pool 获取 AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest
func GetAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest() *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest {
	return poolAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest.Get().(*AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest)
}

// ReleaseAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest 将 AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest(v *AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest) {
	v.Reset()
	poolAlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest.Put(v)
}
