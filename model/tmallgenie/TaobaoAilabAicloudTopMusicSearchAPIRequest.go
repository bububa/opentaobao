package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMusicSearchAPIRequest
对外音乐搜索服务 API请求
taobao.ailab.aicloud.top.music.search

供厂商获取音乐列表 */
type TaobaoAilabAicloudTopMusicSearchAPIRequest struct {
	model.Params
	// botId值
	_botId int64
	// 筛选条件，目前只支持name、type和style
	_params string
	// 分页页码
	_pageNo int64
	// 分页页大小
	_pageSize int64
}

// New
