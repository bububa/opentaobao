package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiMvMusiclistGetAPIRequest
乐馆mv列表 API请求
alibaba.xiami.api.mv.musiclist.get

乐馆mv列表 */
type AlibabaXiamiApiMvMusiclistGetAPIRequest struct {
	model.Params
	// 语种, 有all, chinese, musician, english, japanese, korea
	_type string
	// 分组, 有recommend、hot、new
	_order string
	// 每页记录
	_limit int64
	// 当前页
	_page int64
}

// New
