package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外音乐搜索服务 API请求
taobao.ailab.aicloud.top.music.search

供厂商获取音乐列表
*/
type TaobaoAilabAicloudTopMusicSearchAPIRequest struct {
    model.Params
    // botId值
    _botId   int64
    // 筛选条件，目前只支持name、type和style
    _params   string
    // 分页页码
    _pageNo   int64
    // 分页页大小
    _pageSize   int64
}

// 初始化TaobaoAilabAicloudTopMusicSearchAPIRequest对象
func NewTaobaoAilabAicloudTopMusicSearchRequest() *TaobaoAilabAicloudTopMusicSearchAPIRequest{
    return &TaobaoAilabAicloudTopMusicSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.music.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotId Setter
// botId值
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetBotId(_botId int64) error {
    r._botId = _botId
    r.Set("bot_id", _botId)
    return nil
}

// BotId Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetBotId() int64 {
    return r._botId
}
// Params Setter
// 筛选条件，目前只支持name、type和style
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetParams(_params string) error {
    r._params = _params
    r.Set("params", _params)
    return nil
}

// Params Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetParams() string {
    return r._params
}
// PageNo Setter
// 分页页码
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页页大小
func (r *TaobaoAilabAicloudTopMusicSearchAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAilabAicloudTopMusicSearchAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
