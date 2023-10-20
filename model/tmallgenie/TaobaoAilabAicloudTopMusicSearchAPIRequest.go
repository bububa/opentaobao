package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMusicSearchAPIRequest 对外音乐搜索服务 API请求
// taobao.ailab.aicloud.top.music.search
//
// 供厂商获取音乐列表
type TaobaoAilabAicloudTopMusicSearchAPIRequest struct {
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

// NewTaobaoAilabAicloudTopMusicSearchRequest 初始化TaobaoAilabAicloudTopMusicSearchAPIRequest对象
func NewTaobaoAilabAicloudTopMusicSearchRequest() *TaobaoAilabAicloudTopMusicSearchAPIRequest {
	return &TaobaoAilabAicloudTopMusicSearchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) Reset() {
	r._params = ""
	r._botId = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.music.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// 筛选条件，目前只支持name、type和style
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetParams(_params string) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetParams() string {
	return r._params
}

// SetBotId is BotId Setter
// botId值
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetBotId(_botId int64) error {
	r._botId = _botId
	r.Set("bot_id", _botId)
	return nil
}

// GetBotId BotId Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetBotId() int64 {
	return r._botId
}

// SetPageNo is PageNo Setter
// 分页页码
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页页大小
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoAilabAicloudTopMusicSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMusicSearchRequest()
	},
}

// GetTaobaoAilabAicloudTopMusicSearchRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMusicSearchAPIRequest
func GetTaobaoAilabAicloudTopMusicSearchAPIRequest() *TaobaoAilabAicloudTopMusicSearchAPIRequest {
	return poolTaobaoAilabAicloudTopMusicSearchAPIRequest.Get().(*TaobaoAilabAicloudTopMusicSearchAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMusicSearchAPIRequest 将 TaobaoAilabAicloudTopMusicSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMusicSearchAPIRequest(v *TaobaoAilabAicloudTopMusicSearchAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMusicSearchAPIRequest.Put(v)
}
