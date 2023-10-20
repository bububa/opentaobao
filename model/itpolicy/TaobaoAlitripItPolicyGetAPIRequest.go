package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitpolicygetAPIRequest 【国际机票销售规则】单条查询 API请求
// taobao.alitrip.it.policy.get
//
// 通过此接口可以查询单条销售规则的详情，可以根据taobaoId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据。taobaoId为新增成功时候返回的唯一id，outId为新增时的policy_id（产品编号）
type TaobaoalitripitpolicygetAPIRequest struct {
	model.Params
	// 预留扩展字段
	_extendAttributes string
	// 接入方产品编号
	_outId string
	// 淘宝政策id
	_taobaoId int64
}

// NewTaobaoalitripitpolicygetRequest 初始化TaobaoalitripitpolicygetAPIRequest对象
func NewTaobaoalitripitpolicygetRequest() *TaobaoalitripitpolicygetAPIRequest {
	return &TaobaoalitripitpolicygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripitpolicygetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.policy.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripitpolicygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripitpolicygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// 预留扩展字段
func (r *TaobaoalitripitpolicygetAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extend_attributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoalitripitpolicygetAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetOutId is OutId Setter
// 接入方产品编号
func (r *TaobaoalitripitpolicygetAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("out_id", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoalitripitpolicygetAPIRequest) GetOutId() string {
	return r._outId
}

// SetTaobaoId is TaobaoId Setter
// 淘宝政策id
func (r *TaobaoalitripitpolicygetAPIRequest) SetTaobaoId(_taobaoId int64) error {
	r._taobaoId = _taobaoId
	r.Set("taobao_id", _taobaoId)
	return nil
}

// GetTaobaoId TaobaoId Getter
func (r TaobaoalitripitpolicygetAPIRequest) GetTaobaoId() int64 {
	return r._taobaoId
}
