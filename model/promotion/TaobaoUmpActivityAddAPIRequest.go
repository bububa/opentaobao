package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivityAddAPIRequest
新增优惠活动 API请求
taobao.ump.activity.add

新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家） */
type TaobaoUmpActivityAddAPIRequest struct {
	model.Params
	// 工具id
	_toolId int64
	// 活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。
	_content string
}

// New
