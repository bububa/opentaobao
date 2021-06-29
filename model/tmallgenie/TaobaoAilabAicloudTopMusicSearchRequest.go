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
type TaobaoAilabAicloudTopMusicSearchRequest struct {
    model.Params
    // botId值
    botId   int64
    // 筛选条件，目前只支持name、type和style
    params   string
    // 分页页码
    pageNo   int64
    // 分页页大小
    pageSize   int64
}

// 初始化TaobaoAilabAicloudTopMusicSearchRequest对象
func NewTaobaoAilabAicloudTopMusicSearchRequest() *TaobaoAilabAicloudTopMusicSearchRequest{
    return &TaobaoAilabAicloudTopMusicSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMusicSearchRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.music.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMusicSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BotId Setter
// botId值
func (r *TaobaoAilabAicloudTopMusicSearchRequest) SetBotId(botId int64) error {
    r.botId = botId
    r.Set("bot_id", botId)
    return nil
}

// BotId Getter
func (r TaobaoAilabAicloudTopMusicSearchRequest) GetBotId() int64 {
    return r.botId
}
// Params Setter
// 筛选条件，目前只支持name、type和style
func (r *TaobaoAilabAicloudTopMusicSearchRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r TaobaoAilabAicloudTopMusicSearchRequest) GetParams() string {
    return r.params
}
// PageNo Setter
// 分页页码
func (r *TaobaoAilabAicloudTopMusicSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoAilabAicloudTopMusicSearchRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 分页页大小
func (r *TaobaoAilabAicloudTopMusicSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAilabAicloudTopMusicSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
