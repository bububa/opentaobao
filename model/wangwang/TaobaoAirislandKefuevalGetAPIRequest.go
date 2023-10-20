package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAirislandKefuevalGetAPIRequest 客服评价详情接口_V2 API请求
// taobao.airisland.kefueval.get
//
// 获取买家对客服的服务评价
//
// 注意：
//
// 1. 请求超时[isp.top-remote-connection-timeout]或者数据过大错误[isp.runtime-max-limit]：因为某些帐号请求的数据会非常大，【需要通过减少请求时间范围避免该问题】
//
// 2. 时间范围：[now()-90d&lt;=btime ~ etime &lt;= now()-1d ] AND etime-btime &lt;=7d
//
// 3. 变更eval_recer：可空，返回脱敏的买家nick，如：摩天轮 -&gt; 摩**
//
// 4. 新增labelName：可空
type TaobaoAirislandKefuevalGetAPIRequest struct {
	model.Params
	// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
	_queryIds string
	// 开始时间，格式yyyyMMddHHmmss，时间范围：[now-90d,now-1d]
	_btime string
	// 结束时间，格式yyyyMMddHHmmss，时间范围：[now-90d,now-1d]
	_etime string
}

// NewTaobaoAirislandKefuevalGetRequest 初始化TaobaoAirislandKefuevalGetAPIRequest对象
func NewTaobaoAirislandKefuevalGetRequest() *TaobaoAirislandKefuevalGetAPIRequest {
	return &TaobaoAirislandKefuevalGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAirislandKefuevalGetAPIRequest) GetApiMethodName() string {
	return "taobao.airisland.kefueval.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAirislandKefuevalGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAirislandKefuevalGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryIds is QueryIds Setter
// 客服的nick，多个用逗号分隔，不要超过10个，带cntaobao的长nick
func (r *TaobaoAirislandKefuevalGetAPIRequest) SetQueryIds(_queryIds string) error {
	r._queryIds = _queryIds
	r.Set("query_ids", _queryIds)
	return nil
}

// GetQueryIds QueryIds Getter
func (r TaobaoAirislandKefuevalGetAPIRequest) GetQueryIds() string {
	return r._queryIds
}

// SetBtime is Btime Setter
// 开始时间，格式yyyyMMddHHmmss，时间范围：[now-90d,now-1d]
func (r *TaobaoAirislandKefuevalGetAPIRequest) SetBtime(_btime string) error {
	r._btime = _btime
	r.Set("btime", _btime)
	return nil
}

// GetBtime Btime Getter
func (r TaobaoAirislandKefuevalGetAPIRequest) GetBtime() string {
	return r._btime
}

// SetEtime is Etime Setter
// 结束时间，格式yyyyMMddHHmmss，时间范围：[now-90d,now-1d]
func (r *TaobaoAirislandKefuevalGetAPIRequest) SetEtime(_etime string) error {
	r._etime = _etime
	r.Set("etime", _etime)
	return nil
}

// GetEtime Etime Getter
func (r TaobaoAirislandKefuevalGetAPIRequest) GetEtime() string {
	return r._etime
}
