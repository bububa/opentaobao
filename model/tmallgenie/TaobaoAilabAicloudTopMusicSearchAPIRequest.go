package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmusicsearchAPIRequest 对外音乐搜索服务 API请求
// taobao.ailab.aicloud.top.music.search
//
// 供厂商获取音乐列表
type TaobaoailabaicloudtopmusicsearchAPIRequest struct {
	model.Params
	// 筛选条件，目前只支持name、type和style
	_params string
	// botId值
	_botId int64
	// 分页页码
	_pageNo int64
	// 分页页大小
	_pageSize int64
}

// NewTaobaoailabaicloudtopmusicsearchRequest 初始化TaobaoailabaicloudtopmusicsearchAPIRequest对象
func NewTaobaoailabaicloudtopmusicsearchRequest() *TaobaoailabaicloudtopmusicsearchAPIRequest {
	return &TaobaoailabaicloudtopmusicsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.music.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// 筛选条件，目前只支持name、type和style
func (r *TaobaoailabaicloudtopmusicsearchAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetParams() string {
	return r._params
}

// SetBotId is BotId Setter
// botId值
func (r *TaobaoailabaicloudtopmusicsearchAPIRequest) SetBotId(_botId int64) error {
	r._botId = _botId
	r.Set("bot_id", _botId)
	return nil
}

// GetBotId BotId Getter
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetBotId() int64 {
	return r._botId
}

// SetPageNo is PageNo Setter
// 分页页码
func (r *TaobaoailabaicloudtopmusicsearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页页大小
func (r *TaobaoailabaicloudtopmusicsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoailabaicloudtopmusicsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
