package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopendatarelationqueryAPIRequest 数据关联关系查询 API请求
// alibaba.wdk.marketing.open.data.relation.query
//
// 数据关联关系查询
type AlibabawdkmarketingopendatarelationqueryAPIRequest struct {
	model.Params
	// 外部数据Id
	_outDataIds []string
	// 数据类型：WDK_MARKET:五道口营销
	_bizCode string
	// 数据子类型：ACTIVITY:营销活动数据
	_subBizCode string
}

// NewAlibabawdkmarketingopendatarelationqueryRequest 初始化AlibabawdkmarketingopendatarelationqueryAPIRequest对象
func NewAlibabawdkmarketingopendatarelationqueryRequest() *AlibabawdkmarketingopendatarelationqueryAPIRequest {
	return &AlibabawdkmarketingopendatarelationqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopendatarelationqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.data.relation.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopendatarelationqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopendatarelationqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutDataIds is OutDataIds Setter
// 外部数据Id
func (r *AlibabawdkmarketingopendatarelationqueryAPIRequest) SetOutDataIds(_outDataIds []string) error {
	r._outDataIds = _outDataIds
	r.Set("out_data_ids", _outDataIds)
	return nil
}

// GetOutDataIds OutDataIds Getter
func (r AlibabawdkmarketingopendatarelationqueryAPIRequest) GetOutDataIds() []string {
	return r._outDataIds
}

// SetBizCode is BizCode Setter
// 数据类型：WDK_MARKET:五道口营销
func (r *AlibabawdkmarketingopendatarelationqueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AlibabawdkmarketingopendatarelationqueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetSubBizCode is SubBizCode Setter
// 数据子类型：ACTIVITY:营销活动数据
func (r *AlibabawdkmarketingopendatarelationqueryAPIRequest) SetSubBizCode(_subBizCode string) error {
	r._subBizCode = _subBizCode
	r.Set("sub_biz_code", _subBizCode)
	return nil
}

// GetSubBizCode SubBizCode Getter
func (r AlibabawdkmarketingopendatarelationqueryAPIRequest) GetSubBizCode() string {
	return r._subBizCode
}
